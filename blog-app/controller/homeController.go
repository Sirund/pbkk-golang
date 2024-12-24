package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Welcome to Golang Blog",
	})
}

// import "github.com/gofiber/fiber/v2"

// func Home(c *fiber.Ctx) error {
// 	return c.Render("index", fiber.Map{
// 		"Title": "Welcome to Blog App",
// 	})
// }
