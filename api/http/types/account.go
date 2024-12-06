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

type VerifyOTPRequest struct {
	Email     string    `json:"email" validate:"required,email"`
	Code      string    `json:"code" validate:"required"`
	SessionId uuid.UUID `json:"session_id"`
}

type VerifyOTPResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
