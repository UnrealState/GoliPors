package helpers

import (
	"crypto/rand"
	"github.com/gofiber/fiber/v2"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"golipors/pkg/jwt"
	"regexp"
)

func UserClaims(ctx *fiber.Ctx) *jwt.UserClaims {
	if u := ctx.Locals("user"); u != nil {
		userClaims, ok := u.(*jwt2.Token).Claims.(*jwt.UserClaims)
		if ok {
			return userClaims
		}
	}
	return nil
}

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
