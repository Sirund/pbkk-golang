package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// package controller

// import "github.com/gofiber/fiber/v2"

// func RegisterPage(c *fiber.Ctx) error {
// 	// Render the login page (you can adjust the path to your HTML file)
// 	return c.Render("signup", nil)
// }
