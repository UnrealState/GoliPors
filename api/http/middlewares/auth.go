package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"golipors/internal/user/domain"
	"golipors/internal/user/port"
	"golipors/pkg/jwt"
	"strings"
)

// AuthorizationWithRBAC middleware using jwtware
func AuthorizationWithRBAC(secret []byte, userService port.Service, requiredRoles ...string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secret},
		Claims:      &jwt.UserClaims{},      // Use your custom claims struct
		TokenLookup: "header:Authorization", // Parse JWT from the Authorization header
		AuthScheme:  "Bearer",               // Expect tokens to be prefixed with "Bearer"
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// Extract user claims from the context set by jwtware
			userClaims, ok := ctx.Locals("user").(*jwt.UserClaims)
			if !ok || userClaims == nil {
				return fiber.ErrUnauthorized
			}

			// Fetch the user's details from the database
			user, err := userService.GetUserByID(ctx.Context(), domain.UserID(userClaims.UserID))
			if err != nil {
				return fiber.ErrUnauthorized
			}

			// Check if the user's role matches the required roles
			if !isRoleAllowed(user.Role, requiredRoles) {
				return fiber.ErrForbidden
			}

			// Pass control to the next middleware or handler
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
		},
	})
}

// Helper function to validate roles
func isRoleAllowed(userRole string, allowedRoles []string) bool {
	if len(allowedRoles) == 0 {
		// No role restriction, allow all
		return true
	}
	for _, role := range allowedRoles {
		if strings.EqualFold(userRole, role) {
			return true
		}
	}
	return false
}
