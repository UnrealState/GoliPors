package helpers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ServiceGetter[T any] func(context.Context) T

var validate = validator.New()

func ValidateRequestBody[T any](body T) map[string]string {
	if errs := validate.Struct(body); errs != nil {
		validationErrors := make(map[string]string)

		// Check if the error is of type `validator.ValidationErrors`
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("Validation failed on '%s' with tag '%s'", err.Field(), err.Tag())
		}

		return validationErrors
	}

	return nil
}

func ParseRequestBody[T any](c *fiber.Ctx, body *T) fiber.Map {
	errParse := c.BodyParser(body)
	fmt.Printf("%v\n", body)
	msg := fiber.Map{"error": ErrRequiredBodyNotFound.Error()}

	if errParse != nil {
		msg["message"] = errParse.Error()
		return msg
	}

	errValidation := ValidateRequestBody[T](*body)

	if errValidation != nil {
		msg["details"] = errValidation

		return msg
	}

	return nil
}

func IsValidDate(date string) (time.Time, error) {
	const layout = "2006-01-02" // Reference layout for YYYY-MM-DD
	return time.Parse(layout, date)
}
