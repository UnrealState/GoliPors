package types

import "github.com/google/uuid"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Code      string    `json:"code"`
	SessionId uuid.UUID `json:"session_id"`
}
