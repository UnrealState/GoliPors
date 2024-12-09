package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golipors/api/http/handlers/helpers"
	"golipors/api/http/handlers/presenter"
	"golipors/api/http/types"
	"golipors/app"
	"golipors/config"
	"golipors/pkg/adapters/email"
	"golipors/pkg/cache"

	userService "golipors/internal/user"
	userPort "golipors/internal/user/port"
	jwt2 "golipors/pkg/jwt"
)

type AccountService struct {
	svc                              userPort.Service
	authCache                        *cache.ObjectCache[*presenter.LoginCacheSession]
	emailService                     email.Adapter
	authSecret                       string
	expMin, refreshExpMin, otpTtlMin uint
}

var (
	ErrUserOnCreate      = userService.ErrUserOnCreate
	ErrUserNotFound      = userService.ErrUserNotFound
	ErrUserAlreadyExists = userService.ErrUserAlreadyExists
	ErrCreatingToken     = errors.New("cannot create token")
	ErrBirthdayInvalid   = errors.New("birthday is invalid")
)

func NewAccountService(
	svc userPort.Service,
	cacheService cache.Provider,
	emailService email.Adapter,
	authSecret string,
	expMin, refreshExpMin, otpTtlMin uint,
) *AccountService {
	return &AccountService{
		svc:           svc,
		authCache:     cache.NewJsonObjectCache[*presenter.LoginCacheSession](cacheService, "auth."),
		emailService:  emailService,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
		otpTtlMin:     otpTtlMin,
	}
}

func AccountServiceGetter(appContainer app.App, cfg config.ServerConfig) helpers.ServiceGetter[*AccountService] {
	return func(ctx context.Context) *AccountService {
		return NewAccountService(
			appContainer.UserService(ctx),
			appContainer.Cache(),
			appContainer.MailService(),
			cfg.Secret,
			cfg.AuthExpirationMinutes,
			cfg.AuthRefreshMinutes,
			cfg.OtpTtlMinutes,
		)
	}
}

func (as *AccountService) Login(c context.Context, req types.LoginRequest) (*types.LoginResponse, error) {
	user, err := as.svc.GetUserByUsernamePassword(c, req.Email, req.Password)

	if err != nil {
		return nil, err
	}

	code, err := helpers.GenerateOTP()

	if err != nil {
		return nil, errors.New("error generating OTP")
	}

	log.Println("OTP sent for user", user.ID, "code:", code)

	err = as.emailService.SendText(
		req.Email,
		fmt.Sprintf("GoliPors OTP code for %s", req.Email),
		fmt.Sprintf("GoliPors OTP code: %s", code),
	)

	reqUUID := uuid.New()

	if err != nil {
		log.Println("Error while sending otp:", err)
	}

	err = as.authCache.Set(
		c, strconv.Itoa(int(user.ID)),
		time.Minute*time.Duration(as.otpTtlMin),
		&presenter.LoginCacheSession{
			SessionID: reqUUID,
			UserID:    user.ID,
			Code:      code,
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Code:      code,
		SessionId: reqUUID,
	}, nil
}

func (as *AccountService) VerifyOtp(c context.Context, req types.VerifyOTPRequest) (*types.VerifyOTPResponse, error) {
	user, err := as.svc.GetUserByEmail(c, req.Email)

	if err != nil {
		return &types.VerifyOTPResponse{}, err
	}

	authSession, err := as.authCache.Get(c, strconv.Itoa(int(user.ID)))

	if err != nil {
		return nil, err
	}

	if authSession == nil ||
		authSession.UserID <= 0 ||
		authSession.Code != req.Code ||
		authSession.SessionID != req.SessionId ||
		authSession.UserID != user.ID {
		return nil, ErrUserNotFound
	}

	err = as.authCache.Del(c, strconv.Itoa(int(user.ID)))

	if err != nil {
		return nil, err
	}

	var (
		authExp    = time.Now().Add(time.Minute * time.Duration(as.expMin))
		refreshExp = time.Now().Add(time.Minute * time.Duration(as.refreshExpMin))
	)

	accessToken, err := jwt2.CreateToken([]byte(as.authSecret), jwt2.GenerateUserClaims(user, authExp))
	refreshToken, err := jwt2.CreateToken([]byte(as.authSecret), jwt2.GenerateUserClaims(user, refreshExp))

	if err != nil {
		return nil, ErrCreatingToken
	}

	return &types.VerifyOTPResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *AccountService) Register(c context.Context, req types.RegisterRequest) error {
	newU, err := presenter.RegisterRequestToUserDomain(req)

	if err != nil {
		return ErrBirthdayInvalid
	}

	_, err = as.svc.CreateUser(c, newU)

	if err != nil {
		return err
	}

	return nil
}
