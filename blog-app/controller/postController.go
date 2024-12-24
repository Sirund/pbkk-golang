package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirund/blog-app/database"
	"github.com/sirund/blog-app/models"
	util "github.com/sirund/blog-app/utils"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context) {
	var blogpost models.Blog
	if err := c.ShouldBindJSON(&blogpost); err != nil {
		fmt.Println("Error parsing request body, cause", err)
		c.JSON(400, gin.H{
			"message": "Invalid Payload",
		})
		return
	}

	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Payload",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post successfully created",
	})
}

func AllPost(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	c.JSON(200, gin.H{
		"data": getblog,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func DetailedPost(c *gin.Context) {
	var blogpost models.Blog
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	c.JSON(200, gin.H{
		"data": blogpost,
	})
}

func UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		fmt.Println("Error parsing request body, cause", err)
		c.JSON(400, gin.H{
			"message": "Invalid Payload",
		})
		return
	}

	if err := database.DB.Model(&blog).Updates(blog).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to update post",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post updated successfully",
	})
}

func UniquePost(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	id, _ := util.ParseJwt(cookie)

	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	c.JSON(200, blog)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": "Failed to delete post, record not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}

// package controller

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/sirund/blog-app/database"
// 	"github.com/sirund/blog-app/models"
// 	util "github.com/sirund/blog-app/utils"
// 	"gorm.io/gorm"
// )

// func CreatePost(c *fiber.Ctx) error {
// 	var blogpost models.Blog
// 	if err := c.BodyParser(&blogpost); err != nil {
// 		fmt.Println("Error parsing request body, cause", err)
// 	}

// 	if err := database.DB.Create(&blogpost).Error; err != nil {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Invalid Payload",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Post successfully created",
// 	})
// }

// func AllPost(c *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(c.Query("page", "1"))
// 	limit := 5
// 	offset := (page - 1) * limit
// 	var total int64
// 	var getblog []models.Blog
// 	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
// 	database.DB.Model(&models.Blog{}).Count(&total)
// 	return c.JSON(fiber.Map{
// 		"data": getblog,
// 		"meta": fiber.Map{
// 			"total":     total,
// 			"page":      page,
// 			"last_page": math.Ceil(float64(int(total) / limit)),
// 		},
// 	})
// }

// func DetailedPost(c *fiber.Ctx) error {
// 	var blogpost models.Blog
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
// 	return c.JSON(fiber.Map{
// 		"data": blogpost,
// 	})
// }

// func UpdatePost(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	blog := models.Blog{
// 		Id: uint(id),
// 	}

// 	if err := c.BodyParser(&blog); err != nil {
// 		fmt.Println("Error parsing request body, cause", err)
// 	}
// 	database.DB.Model(&blog).Updates(blog)
// 	return c.JSON(fiber.Map{
// 		"message": "Post upadated successfully",
// 	})
// }

// func UniquePost(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")
// 	id, _ := util.ParseJwt(cookie)

// 	var blog []models.Blog
// 	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
// 	return c.JSON(blog)
// }

// func DeletePost(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	blog := models.Blog{
// 		Id: uint(id),
// 	}

// 	deleteQuery := database.DB.Delete(&blog)
// 	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Failed to delete post, record not found",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Post deleted successfully",
// 	})
// }
