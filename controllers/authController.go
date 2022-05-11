package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"restAuth/database"
	"restAuth/models"
	"time"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Register(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "Error hashing password",
			"error":  err.Error(),
		})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(hashedPassword),
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"status": "success",
		"user":   user,
	})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"status": err.Error(),
		})
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status": "User not found",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status": "Wrong password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    string(user.ID),
	})

	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status": "Error signing token",
		})
	}

	cookie := &fiber.Cookie{
		Name:     "jwt-token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"status": "logged in",
		"user":   user,
	})

}
