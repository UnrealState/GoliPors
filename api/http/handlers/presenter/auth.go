package presenter

import (
	"github.com/google/uuid"
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
