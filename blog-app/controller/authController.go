package controller

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirund/blog-app/database"
	"github.com/sirund/blog-app/models"
	util "github.com/sirund/blog-app/utils"
)

func Register(c *gin.Context) {
	var data map[string]interface{}
	var userData models.User

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	// Check if the password is less than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password Must Be Greater Than 6 Characters",
		})
		return
	}

	// Check if the email is valid
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Email Address",
		})
		return
	}

	// Check if the email already exists in database
	database.DB.Where("email = ?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email Already Exists",
		})
		return
	}

	// Add user to database
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	if err := database.DB.Create(&user).Error; err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "User Created Successfully",
	})
}

func validateEmail(email string) bool {
	validate := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return validate.MatchString(email)
}

func Login(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Email Address Doesn't Exist",
		})
		return
	}

	if err := user.CheckPassword(data["password"]); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Password",
		})
		return
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})
		return
	}

	// Set JWT as a cookie
	c.SetCookie("jwt", token, int(24*time.Hour.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "Login Successful",
	})
}

// package controller

// import (
// 	"fmt"
// 	"log"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/sirund/blog-app/database"
// 	"github.com/sirund/blog-app/models"
// 	util "github.com/sirund/blog-app/utils"
// )

// func Register(c *fiber.Ctx) error {
// 	var data map[string]interface{}
// 	var userData models.User
// 	if err := c.BodyParser(&data); err != nil {
// 		fmt.Println("Error parsing request body, cause", err)
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Invalid request body",
// 		})
// 	}

// 	// Check if the password less than 6 characters
// 	if len(data["password"].(string)) <= 6 {
// 		c.Status(400) // Bad request
// 		return c.JSON(fiber.Map{
// 			"message": "Password Must Be Greather Than 6 Characters",
// 		})
// 	}

// 	// Check if the email is valid
// 	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Invalid Email Address",
// 		})
// 	}

// 	// Check if the email already exists in database
// 	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
// 	if userData.Id != 0 {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Email Already Exists",
// 		})
// 	}

// 	// Add data of user to database
// 	user := models.User{
// 		FirstName: data["first_name"].(string),
// 		LastName:  data["last_name"].(string),
// 		Phone:     data["phone"].(string),
// 		Email:     strings.TrimSpace(data["email"].(string)),
// 	}

// 	user.SetPassword(data["password"].(string))
// 	if err := database.DB.Create(&user); err != nil {
// 		log.Print(err)
// 	}

// 	c.Status(200)
// 	return c.JSON(fiber.Map{
// 		"user":    user,
// 		"message": "User Created Successfully",
// 	})

// }

// func validateEmail(email string) bool {
// 	validate := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
// 	return validate.MatchString(email)
// }

// func Login(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Invalid request body",
// 		})
// 	}

// 	var user models.User
// 	database.DB.Where("email=?", data["email"]).First(&user)
// 	if user.Id == 0 {
// 		c.Status(404)
// 		return c.JSON(fiber.Map{
// 			"message": "Email Address Doesn't Exist",
// 		})
// 	}

// 	if err := user.CheckPassword(data["password"]); err != nil {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Invalid Password",
// 		})
// 	}

// 	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
// 	if err != nil {
// 		c.Status(fiber.StatusInternalServerError)
// 		return nil
// 	}

// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    token,
// 		Expires:  time.Now().Add(time.Hour * 24),
// 		HTTPOnly: true,
// 	}

// 	c.Cookie(&cookie)
// 	return c.JSON(fiber.Map{
// 		"user":    user,
// 		"message": "Login Successful",
// 	})
// }

// type Claims struct {
// 	jwt.StandardClaims
// }
