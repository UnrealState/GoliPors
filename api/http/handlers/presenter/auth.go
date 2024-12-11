package presenter

import (
	"errors"
	"github.com/google/uuid"
	"golipors/api/http/handlers/helpers"
	"golipors/api/http/types"
	"golipors/internal/user/domain"
)

type UserToken struct {
	AuthorizationToken string
	RefreshToken       string
	ExpiresAt          int64
}

type LoginCacheSession struct {
	SessionID uuid.UUID
	UserID    domain.UserID
	Code      string
}

func RegisterRequestToUserDomain(req types.RegisterRequest) (*domain.User, error) {
	birthday, err := helpers.IsValidDate(req.Birthday)

	if err != nil {
		return nil, errors.New("birthday invalid")
	}

	return &domain.User{
		NationalID:    req.NationalID,
		Email:         req.Email,
		Password:      req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Birthday:      birthday,
		City:          req.City,
		WalletBalance: 0,
		VoteBalance:   0,
	}, nil
}
