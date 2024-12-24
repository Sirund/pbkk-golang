package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirund/blog-app/database"
	"github.com/sirund/blog-app/routes"
)

func main() {
	// Connect to the database
	database.Connect()

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a Gin router instance
	router := gin.Default()

	// Set up HTML template rendering
	router.LoadHTMLGlob("./templates/*")

	// Serve static files
	router.Static("/static", "./static")

	// Register routes
	routes.Setup(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Fallback port
	}
	router.Run(":" + port)
}

// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/template/html/v2"
// 	"github.com/joho/godotenv"
// 	"github.com/sirund/blog-app/database"
// 	"github.com/sirund/blog-app/routes"
// )

// func main() {
// 	database.Connect()
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	engine := html.New("./templates", ".html")
// 	app := fiber.New(fiber.Config{
// 		Views: engine,
// 	})

// 	app.Static(
// 		"/static",
// 		"./static",
// 	)

// 	routes.Setup(app)
// 	port := os.Getenv("PORT")
// 	app.Listen(":" + port)

// }
