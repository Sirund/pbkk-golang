package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randWord(n int) string {
	word := make([]rune, n)
	for i := range word {
		word[i] = letters[rand.Intn(len(letters))]
	}
	return string(word)
}

func UploadImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	files := form.File["image"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	fileName := ""
	for _, file := range files {
		// Generate a random file name
		fileName = randWord(5) + "-" + filepath.Base(file.Filename)

		// Save the file to the specified directory
		if err := c.SaveUploadedFile(file, "./images/"+fileName); err != nil {
			fmt.Println("Failed to save image:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
	}

	// Return the URL of the uploaded image
	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:3000/api/uploads/" + fileName,
	})
}

// package controller

// import (
// 	"fmt"
// 	"math/rand"

// 	"github.com/gofiber/fiber/v2"
// )

// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randWord(n int) string {
// 	word := make([]rune, n)
// 	for i := range word {
// 		word[i] = letters[rand.Intn(len(letters))]
// 	}

// 	return string(word)
// }

// func UploadImage(c *fiber.Ctx) error {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return err
// 	}

// 	files := form.File["image"]
// 	fileName := ""

// 	for _, file := range files {
// 		fileName = randWord(5) + "-" + file.Filename
// 		if err := c.SaveFile(file, "./images/"+fileName); err != nil {
// 			fmt.Println("Failed to save image")
// 			return nil
// 		}
// 	}

// 	return c.JSON(fiber.Map{
// 		"url": "http://localhost:3000/api/uploads/" + fileName,
// 	})

// }
