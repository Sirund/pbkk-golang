package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/sirund/blog-app/utils"
)

func IsAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT cookie
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		// Parse and validate the JWT
		if _, err := util.ParseJwt(cookie); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		// Proceed to the next middleware or handler
		c.Next()
	}
}

// package middleware

// import (
// 	"github.com/gofiber/fiber/v2"
// 	util "github.com/sirund/blog-app/utils"
// )

// func IsAuthenticate(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")

// 	if _, err := util.ParseJwt(cookie); err != nil {
// 		c.Status(fiber.StatusUnauthorized)
// 		return c.JSON(fiber.Map{
// 			"message": "Unauthorized",
// 		})
// 	}
// 	return c.Next()

// }
