package Controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/Database"
	"github.com/pred695/Go-JWT/Models"
	"github.com/pred695/Go-JWT/Utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	FormData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func GetUsers(ctx *fiber.Ctx) error {
	contextMap := fiber.Map{
		"message":    "Get Users",
		"statusText": "Ok",
	}
	db := Database.DbConn
	var users []Models.User
	result := db.Find(&users)
	if result.Error != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = "Error Fetching Users"
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}
	contextMap["users"] = users
	return ctx.JSON(contextMap)
}

func LoginUser(ctx *fiber.Ctx) error {
	contextMap := fiber.Map{
		"message":    "User Logged In",
		"statusText": "Token generated successfully", /*stored in cookie*/
	}

	db := Database.DbConn

	var loginCredentials LoginData
	if err := ctx.BodyParser(&loginCredentials); err != nil {
		contextMap["statusText"] = "Bad Request"
		contextMap["message"] = "Ok"
		return ctx.Status(fiber.StatusBadRequest).JSON(contextMap)
	}
	var user Models.User

	//search the user in the database according to the given email.
	db.First(&user, "username = ?", loginCredentials.Username)
	if user.ID == 0 {
		contextMap["statusText"] = "Not Found"
		contextMap["message"] = "User not found"
		return ctx.Status(fiber.StatusNotFound).JSON(contextMap)
	}

	//if user is found, validate password:

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCredentials.Password))
	if err != nil {
		fmt.Println("Invalid Password")
		contextMap["statusText"] = "Unauthorized"
		contextMap["message"] = "Invalid Password"
		return ctx.Status(fiber.StatusUnauthorized).JSON(contextMap)
	}

	//create token:
	token, err := Utils.GenerateToken(&user)
	if err != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = "Error in generating token"
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}
	contextMap["token"] = token
	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(10 * time.Minute), // Set expiry time as needed
		HTTPOnly: true,                             // HTTPOnly to prevent access via JavaScript
		Secure:   false,                            // Secure flag to ensure cookie is sent over HTTPS
		SameSite: "Strict",                         // SameSite policy to prevent CSRF attacks
	})
	return ctx.Status(fiber.StatusOK).JSON(contextMap)
}

func RegisterUser(ctx *fiber.Ctx) error {
	contextMap := fiber.Map{
		"message":    "Register User",
		"statusText": "Ok",
	}
	db := Database.DbConn
	var signUpCredentials FormData
	user := new(Models.User)
	if err := ctx.BodyParser(&signUpCredentials); err != nil {
		contextMap["statusText"] = "Bad Request"
		contextMap["message"] = err
		return ctx.Status(fiber.StatusBadRequest).JSON(contextMap)
	}
	user.Username = signUpCredentials.Username
	user.Password = Utils.HashPassword(signUpCredentials.Password)
	user.Email = signUpCredentials.Email
	result := db.Create(&user)
	if result.Error != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = "Username or email already exists"
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}
	contextMap["user"] = user
	return ctx.Status(fiber.StatusCreated).JSON(contextMap)
}

func LogOutUser(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour), // Set an expired time
	})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged out successfully",
	})
}

func RefreshToken(ctx *fiber.Ctx) error {

	return nil
}
