package middlewares

import (
	"github.com/casbin/casbin/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"golipors/internal/user/domain"
	"golipors/internal/user/port"
	"golipors/pkg/jwt"
)

// AuthorizationWithRBAC middleware using jwtware
func AuthorizationWithRBAC(secret []byte, userService port.Service, enforcer *casbin.Enforcer, requiredPermissions ...string) fiber.Handler {
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

			// Check if the user has the required permissions using Casbin
			if !hasPermission(enforcer, user.Role, requiredPermissions) {
				return fiber.ErrForbidden
			}

			// Proceed to the next handler if the permission check passes
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
		},
	})
}

// Helper function to check if the user's role has the required permission using Casbin
func hasPermission(enforcer *casbin.Enforcer, userRole string, requiredPermissions []string) bool {
	// Check if the user's role has the required permission in Casbin
	for _, permission := range requiredPermissions {
		allowed, err := enforcer.Enforce(userRole, permission)
		if err != nil || !allowed {
			return false
		}
	}
	return true
}
