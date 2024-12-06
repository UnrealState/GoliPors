package middlerwares

import (
	"github.com/gofiber/fiber/v2"
	"golipors/pkg/context"
	"gorm.io/gorm"
)

func SetUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext()))
	return c.Next()
}

func SetTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tx := db.Begin()

		context.SetDB(c.UserContext(), tx, true)

		err := c.Next()

		if c.Response().StatusCode() >= 300 {
			return context.Rollback(c.UserContext())
		}

		if err := context.CommitOrRollback(c.UserContext(), true); err != nil {
			return err
		}

		return err
	}
}
