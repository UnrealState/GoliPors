package helpers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ServiceGetter[T any] func(context.Context) T

var validate = validator.New()

func ValidateRequestBody[T any](body T) map[string]string {
	if err := validate.Struct(body); err != nil {
		validationErrors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("Validation failed on '%s' with tag '%s'", err.Field(), err.Tag())
		}

		return validationErrors
	}

	return nil
}

func ParseRequestBody[T any](c *fiber.Ctx, body *T) error {
	errParse := c.BodyParser(body)
	errValidation := ValidateRequestBody[*T](body)

	if errParse != nil || errValidation != nil {
		msg := fiber.Map{"error": ErrRequiredBodyNotFound}

		if errValidation != nil {
			msg["details"] = errValidation
		}

		return c.Status(fiber.StatusBadRequest).JSON(msg)
	}

	return nil
}
