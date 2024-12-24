package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// package controller

// import "github.com/gofiber/fiber/v2"

// func LoginPage(c *fiber.Ctx) error {
// 	// Render the login page (you can adjust the path to your HTML file)
// 	return c.Render("login", nil)
// }
