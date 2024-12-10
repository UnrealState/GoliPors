package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golipors/api/http/handlers/helpers"
	middlerwares "golipors/api/http/middlewares"
	"golipors/api/http/services"
	"golipors/app"
	"golipors/config"
)

func RegisterQuestionnaireHandlers(router fiber.Router, appContainer app.App, cfg config.ServerConfig) {
	accountGroup := router.Group("/questionnaire")
	accountSvcGetter := services.QuestionnaireServiceGetter(appContainer, cfg)

	accountGroup.Use(middlerwares.SetTransaction(appContainer.DB()))

	accountGroup.Get("/:id", GetQuestionnaire(accountSvcGetter))
	accountGroup.Post("/", CreateQuestionnaire(accountSvcGetter))
}

func GetQuestionnaire(svcGetter helpers.ServiceGetter[*services.QuestionnaireService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func CreateQuestionnaire(svcGetter helpers.ServiceGetter[*services.QuestionnaireService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
