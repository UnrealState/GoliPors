// api/http/handlers/survey.go
package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"golipors/api/http/handlers/presenter"
	"golipors/api/http/mapper"
	"golipors/internal/survey/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type SurveyHandler struct {
	service port.Service
}

func NewSurveyHandler(service port.Service) *SurveyHandler {
	return &SurveyHandler{
		service: service,
	}
}

func (h *SurveyHandler) RegisterRoutes(api fiber.Router) {
	api.Post("/surveys", h.CreateSurvey)
	api.Get("/surveys/:id", h.GetSurveyByID)
	api.Put("/surveys/:id", h.UpdateSurvey)
	api.Delete("/surveys/:id", h.DeleteSurvey)
}

// CreateSurvey handles POST /api/surveys
func (h *SurveyHandler) CreateSurvey(c *fiber.Ctx) error {
	var req presenter.CreateSurveyRequest

	// Parse the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate request input
	if err := validate.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors[e.Field()] = fmt.Sprintf("Validation failed on '%s' with tag '%s'", e.Field(), e.Tag())
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationErrors,
		})
	}

	// Example: Retrieve ownerID from authenticated user context; here, we use a placeholder
	ownerID := uint(1)

	// Convert from presenter to domain model
	survey := mapper.CreateSurveyRequestToDomain(req, ownerID)

	// Create survey using the service
	surveyID, err := h.service.CreateSurvey(c.Context(), survey)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Build response using presenter
	response := presenter.CreateSurveyResponse{
		ID:      surveyID,
		Title:   survey.Title,
		OwnerID: ownerID,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// GetSurveyByID handles GET /api/surveys/:id
func (h *SurveyHandler) GetSurveyByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid survey ID",
		})
	}

	survey, err := h.service.GetSurveyByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if survey == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Survey not found",
		})
	}

	response := mapper.DomainSurveyToGetSurveyResponse(*survey)
	return c.Status(http.StatusOK).JSON(response)
}

// UpdateSurvey handles PUT /api/surveys/:id
func (h *SurveyHandler) UpdateSurvey(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid survey ID",
		})
	}

	var req presenter.UpdateSurveyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert from presenter to domain model
	survey := mapper.UpdateSurveyRequestToDomain(req)
	survey.ID = uint(id)

	if err := h.service.UpdateSurvey(c.Context(), survey); err != nil {
		if err.Error() == "survey not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Survey not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updatedSurvey, err := h.service.GetSurveyByID(c.Context(), uint(id))
	if err != nil || updatedSurvey == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve updated survey",
		})
	}

	response := mapper.DomainSurveyToUpdateSurveyResponse(*updatedSurvey)
	return c.Status(http.StatusOK).JSON(response)
}

// DeleteSurvey handles DELETE /api/surveys/:id
func (h *SurveyHandler) DeleteSurvey(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid survey ID",
		})
	}

	if err := h.service.DeleteSurvey(c.Context(), uint(id)); err != nil {
		if err.Error() == "survey not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Survey not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusOK)
}
