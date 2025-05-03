package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SachinDuhan/multiplayer/config"
	"github.com/SachinDuhan/multiplayer/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GenerateAccessAndRefreshTokens(user models.Player) (string, string, error) {
  accessToken, err := user.GenerateAccessToken()
  if err != nil {
    return "", "", fiber.NewError(http.StatusInternalServerError, "Failed to generate access token")
  }


  refreshToken, err := user.GenerateRefreshToken()
  if err != nil {
    return "", "", fiber.NewError(http.StatusInternalServerError, "Failed to generate refresh token")
  }

  user.RefreshToken = refreshToken

  db := config.DB

	if err := db.Save(&user).Error; err != nil {
		return "", "", fiber.NewError(http.StatusInternalServerError, "Failed to save refresh token")
	}

	return accessToken, refreshToken, nil
}

// RegisterUser handles user registration
func RegisterUser(c *fiber.Ctx) error {
  // Define a request struct for validation
  type RegisterRequest struct {
    Name string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required,min=6"`
  }

  var req RegisterRequest
  if err := c.BodyParser(&req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
  }

  // Create user model
  user := models.Player{
    Name: req.Name,
    Email: req.Email,
    Password: req.Password,
  }

  // Save to DB
  db := config.DB
  if err := db.Create(&user).Error; err != nil {
    log.Println("Error creating user:", err)
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not register user"})
  }

  return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
  type LoginRequest struct {
    Name string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required,min=6"`
  }
  
  var req LoginRequest
  if err := c.BodyParser(&req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
  }

  if req.Name == "" && req.Email == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Either Name or Email is needed"})
  }

  db := config.DB

  var user models.Player
  result := db.Where("Name = ? OR Email = ?", req.Name, req.Email).First(&user)

  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound{
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
    }
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error})
  }

  isPasswordValid := user.CheckPassword(req.Password)

  if isPasswordValid != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Passord"})
  }

  accessToken, refreshToken, err := GenerateAccessAndRefreshTokens(user)

  if err != nil {
    return fiber.NewError(http.StatusInternalServerError, "Could not generate access and refresh token")
  }

  var loggedInUser models.Player
  if err := db.Select("id, name, email").Where("id = ?", user.ID).First(&loggedInUser).Error; err != nil {
    return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch logged in user")
  }

  cookieOptions := fiber.Cookie{
		HTTPOnly: true,
		Secure:   true,
	}

	// Set cookies
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		HTTPOnly: cookieOptions.HTTPOnly,
		Secure:   cookieOptions.Secure,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HTTPOnly: cookieOptions.HTTPOnly,
		Secure:   cookieOptions.Secure,
	})

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":         loggedInUser,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"message":      "User logged in successfully",
	})
}


func GetPlayerInfo(c *fiber.Ctx) error {
  id := c.Query("id") 

  if id == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error": "Player ID is required",
    })
  }

  db := config.DB

  var user models.Player
  fmt.Print(id)
  result := db.Select("ID, Name, Email").Where("ID = ?", id).First(&user)

  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound{
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
    }
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error})
  }

  return c.Status(fiber.StatusCreated).JSON(user)
}
