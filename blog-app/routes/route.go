package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirund/blog-app/controller"
	"github.com/sirund/blog-app/middleware"
)

func Setup(router *gin.Engine) {
	// Public routes
	router.GET("/", controller.Home)
	router.GET("/login", controller.LoginPage)
	router.GET("/register", controller.RegisterPage)
	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)

	// Static file serving
	router.Static("/api/uploads", "./images")

	// Authenticated routes
	authRoutes := router.Group("/api", middleware.IsAuthenticate())
	{
		authRoutes.POST("/post", controller.CreatePost)
		authRoutes.POST("/upload-image", controller.UploadImage)
		authRoutes.GET("/allpost", controller.AllPost)
		authRoutes.GET("/allpost/:id", controller.DetailedPost)
		authRoutes.PUT("/updatepost/:id", controller.UpdatePost)
		authRoutes.GET("/uniquepost", controller.UniquePost)
		authRoutes.DELETE("/deletepost/:id", controller.DeletePost)
	}
}

// package routes

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/sirund/blog-app/controller"
// 	"github.com/sirund/blog-app/middleware"
// )

// func Setup(app *fiber.App) {
// 	app.Get("/", controller.Home)
// 	app.Get("/login", controller.LoginPage)
// 	app.Get("/register", controller.RegisterPage)
// 	app.Post("/api/register", controller.Register)
// 	app.Post("/api/login", controller.Login)

// 	app.Use(middleware.IsAuthenticate)
// 	app.Post("/api/post", controller.CreatePost)
// 	app.Post("/api/upload-image", controller.UploadImage)
// 	app.Get("/api/allpost", controller.AllPost)
// 	app.Get("/api/allpost/:id", controller.DetailedPost)
// 	app.Put("/api/updatepost/:id", controller.UpdatePost)
// 	app.Get("/api/uniquepost", controller.UniquePost)
// 	app.Delete("/api/deletepost/:id", controller.DeletePost)
// 	app.Static("/api/uploads", "./images")
// }
