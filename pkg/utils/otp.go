package utils

import (
	"crypto/rand"
	"fmt"
	"golipors/pkg/redis"
	"time"
)

func GenerateOTP() (string, error) {
	n := 6
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < n; i++ {
		bytes[i] = (bytes[i] % 10) + '0'
	}

	return string(bytes), nil
}

func SendOTPEmail(email, otp string) error {
	// Implement email sending logic here
	fmt.Printf("Sending OTP %s to email %s\n", otp, email)
	return nil
}

func StoreOTP(email, otp string) error {
	key := fmt.Sprintf("otp:%s", email)
	return redis.Client.Set(key, otp, 5*time.Minute).Err()
}

func ValidateOTP(email, otp string) (bool, error) {
	key := fmt.Sprintf("otp:%s", email)
	storedOTP, err := redis.Client.Get(key).Result()
	if err != nil {
		return false, err
	}
	return storedOTP == otp, nil
}
