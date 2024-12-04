package presenter

import (
	"github.com/google/uuid"
	"golipors/internal/user/domain"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

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
