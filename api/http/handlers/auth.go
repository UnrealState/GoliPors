// api/http/handlers/auth.go
package handlers

import (
	"golipors/pkg/models"
	"golipors/pkg/postgres"
	"golipors/pkg/utils"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RegisterInput struct {
	NationalID  string `json:"national_id" validate:"required,len=10"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"` // Format: YYYY-MM-DD
	City        string `json:"city"`
}

func Register(c *fiber.Ctx) error {
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error:": err.Error()})
	}

	// Validate National ID
	if !isValidIranianNationalID(input.NationalID) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid national ID"})
	}

	// Parse Date of Birth
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date of birth format"})
	}

	user := models.User{
		NationalID:       input.NationalID,
		Email:            strings.ToLower(input.Email),
		Password:         input.Password, // Will be hashed by BeforeCreate hook
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		DateOfBirth:      dob,
		RegistrationDate: time.Now(),
		City:             input.City,
	}

	// Save user to database
	result := postgres.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Registration successful"})
}

func isValidIranianNationalID(id string) bool {
	if len(id) != 10 {
		return false
	}
	matched, _ := regexp.MatchString(`^\d{10}$`, id)
	return matched
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var user models.User
	result := postgres.DB.Where("email = ?", strings.ToLower(input.Email)).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Generate OTP
	otp, err := utils.GenerateOTP()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate OTP"})
	}

	// Store OTP in Redis
	err = utils.StoreOTP(user.Email, otp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not store OTP"})
	}

	// Send OTP via Email
	err = utils.SendOTPEmail(user.Email, otp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not send OTP"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OTP sent to your email"})
}

type VerifyOTPInput struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required"`
}

func VerifyOTP(c *fiber.Ctx) error {
	var input VerifyOTPInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	valid, err := utils.ValidateOTP(input.Email, input.OTP)
	if err != nil || !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OTP"})
	}

	var user models.User
	result := postgres.DB.Where("email = ?", strings.ToLower(input.Email)).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email"})
	}

	// Generate JWT Token
	token, err := utils.GenerateToken(user.ID, 60) // Expires in 60 minutes
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

type ResetPasswordRequestInput struct {
	Email string `json:"email" validate:"required,email"`
}

func RequestPasswordReset(c *fiber.Ctx) error {
	var input ResetPasswordRequestInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if user exists
	var user models.User
	result := postgres.DB.Where("email = ?", strings.ToLower(input.Email)).First(&user)
	if result.Error != nil {
		// To prevent email enumeration, always return the same response
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "If the email exists, a reset link will be sent"})
	}

	// Generate reset token
	resetToken, err := utils.GenerateResetToken(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate reset token"})
	}

	// Send reset link via email
	err = utils.SendPasswordResetEmail(user.Email, resetToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not send reset email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "If the email exists, a reset link will be sent"})
}

type ResetPasswordInput struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

func ResetPassword(c *fiber.Ctx) error {
	var input ResetPasswordInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	email, err := utils.ValidateResetToken(input.Token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	var user models.User
	result := postgres.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}

	// Update password
	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	user.Password = hashedPassword
	postgres.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Password reset successful"})
}

func Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
}
