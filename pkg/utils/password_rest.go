package utils

import (
	"fmt"
	"golipors/pkg/redis"
	"time"

	"github.com/google/uuid"
)

// GenerateResetToken generates a unique token for password reset and stores it in Redis.
func GenerateResetToken(email string) (string, error) {
	token := uuid.New().String()
	key := fmt.Sprintf("reset:%s", token)
	err := redis.Client.Set(key, email, 15*time.Minute).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

// ValidateResetToken validates the password reset token and retrieves the associated email.
func ValidateResetToken(token string) (string, error) {
	key := fmt.Sprintf("reset:%s", token)
	email, err := redis.Client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return email, nil
}

// SendPasswordResetEmail sends a password reset email to the user.
func SendPasswordResetEmail(email, token string) error {
	resetLink := fmt.Sprintf("https://yourdomain.com/reset-password?token=%s", token)
	// Implement email sending logic here using an email service provider
	fmt.Printf("Sending password reset link %s to email %s\n", resetLink, email)
	return nil
}
