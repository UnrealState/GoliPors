package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golipors/api/http/handlers/helpers"
	"golipors/api/http/handlers/presenter"
	"golipors/api/http/types"
	"golipors/app"
	"golipors/config"
	userService "golipors/internal/user"
	userPort "golipors/internal/user/port"
	"golipors/pkg/adapters/email"
	"golipors/pkg/cache"
	"log"
	"strconv"
	"time"
)

type AccountService struct {
	svc                              userPort.Service
	cacheService                     cache.Provider
	emailService                     email.Adapter
	authSecret                       string
	expMin, refreshExpMin, otpTtlMin uint
}

var (
	ErrUserCreationValidation = userService.ErrUserCreationValidation
	ErrUserOnCreate           = userService.ErrUserOnCreate
	ErrUserNotFound           = userService.ErrUserNotFound
	ErrInvalidPassword        = userService.ErrInvalidPassword
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
		cacheService:  cacheService,
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
			appContainer.UserService(),
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
		return &types.LoginResponse{}, err
	}

	code, err := helpers.GenerateOTP()

	if err != nil {
		return &types.LoginResponse{}, errors.New("error generating OTP")
	}

	log.Println("OTP sent for user", user.ID, "code:", code)

	oc := cache.NewJsonObjectCache[presenter.LoginCacheSession](as.cacheService, "auth.")

	err = as.emailService.SendText(
		req.Email,
		fmt.Sprintf("GoliPors OTP code for %s", req.Email),
		fmt.Sprintf("GoliPors OTP code: %s", code),
	)

	log.Println(err)

	err = oc.Set(c, strconv.Itoa(int(user.ID)), time.Minute*time.Duration(as.otpTtlMin), presenter.LoginCacheSession{
		SessionID: uuid.New(),
		UserID:    user.ID,
		Code:      code,
	})

	if err != nil {
		return &types.LoginResponse{}, err
	}

	return &types.LoginResponse{
		Code:      code,
		SessionId: uuid.New(),
	}, nil
}
