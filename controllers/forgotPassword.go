package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/smtp"
	"restAuth/database"
	"restAuth/models"
)

func ForgotPassword(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	token := RandomString(12)

	pwReset := models.PasswordReset{
		Token: token,
		Email: data["email"],
	}

	database.DB.Create(&pwReset)

	to := []string{data["email"]}
	from := "admin@gmail.com"

	//subject := "Password reset"
	body := "Click on the link to reset your password: http://localhost:3000/reset/" + token
	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, []byte(body))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error sending email",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"token":   token,
		"message": "Password reset link sent to your email",
	})

}

func ResetPassword(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	var pwReset models.PasswordReset
	database.DB.Where("token = ?", data["token"]).Last(&pwReset)

	if pwReset.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid token",
		})
	}

	var user models.User
	database.DB.Where("email = ?", pwReset.Email).First(&user)

	hash, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	user.Password = string(hash)
	database.DB.Save(&user)

	database.DB.Delete(&pwReset)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Password reset successfully",
	})

}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
