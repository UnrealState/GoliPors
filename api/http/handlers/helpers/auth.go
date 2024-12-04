package helpers

import (
	"crypto/rand"
	"regexp"
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

// IsValidEmail validates the structure of an email address
func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
