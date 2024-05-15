package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/database"
	"github.com/pred695/Go-JWT/models"
	"github.com/pred695/Go-JWT/utils"
)

func ListTasks(ctx *fiber.Ctx) error {
	contextMap := fiber.Map{
		"message":    "List of tasks",
		"statusText": "Ok",
	}
	db := database.DbConn
	var tasks []models.Task

	tokenString := ctx.Cookies("token")
	if tokenString == "" {
		contextMap["statusText"] = "Unauthorized"
		contextMap["message"] = "Token not found"
		return ctx.Status(fiber.StatusUnauthorized).JSON(contextMap)
	}

	//Validate Token if found:
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		contextMap["statusText"] = "Unauthorized"
		contextMap["message"] = "Invalid Token"
		return ctx.Status(fiber.StatusUnauthorized).JSON(contextMap)
	}

	result := db.Where("username = ?", claims.Username).Find(&tasks)
	if result.Error != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = "Error fetching tasks"
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}
	contextMap["tasks"] = tasks
	return ctx.Status(fiber.StatusOK).JSON(contextMap)
}

func CreateTask(ctx *fiber.Ctx) error {
	contextMap := fiber.Map{
		"message":    "Task created",
		"statusText": "Ok",
	}
	db := database.DbConn
	task := new(models.Task)

	if err := ctx.BodyParser(task); err != nil {
		contextMap["statusText"] = "Bad Request"
		contextMap["message"] = "Error parsing the request"
		return ctx.Status(fiber.StatusBadRequest).JSON(contextMap)
	}
	tokenString := ctx.Cookies("token")
	if tokenString == "" {
		contextMap["statusText"] = "Unauthorized"
		contextMap["message"] = "Token not found"
		return ctx.Status(fiber.StatusUnauthorized).JSON(contextMap)
	}

	//Validate Token if found:
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		contextMap["statusText"] = "Unauthorized"
		contextMap["message"] = "Invalid Token"
		return ctx.Status(fiber.StatusUnauthorized).JSON(contextMap)
	}
	var user models.User
	findUser := db.Where("username = ?", claims.Username).First(&user)
	if findUser.Error != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = "Error fetching user"
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}
	task.Username = user.Username
	result := db.Create(&task)
	if result.Error != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = "Error creating a new task"
		ctx.Status(500)
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}
	fmt.Println(result)
	contextMap["task"] = task
	ctx.Status(fiber.StatusCreated).JSON(contextMap)

	return ctx.JSON(contextMap)
}

func UpdateTask(ctx *fiber.Ctx) error {
	contextMap := fiber.Map{
		"message":    "Task updated",
		"statusText": "Ok",
	}
	db := database.DbConn
	taskId := ctx.Params("id")
	var task models.Task

	db.First(&task, taskId) //store the task in the task variable

	if task.ID == 0 {
		contextMap["statusText"] = "Not Found"
		contextMap["message"] = "Task not found"
		return ctx.Status(fiber.StatusNotFound).JSON(contextMap)
	}

	err := ctx.BodyParser(&task)
	if err != nil {
		contextMap["statusText"] = "Bad Request"
		contextMap["message"] = err
		return ctx.Status(fiber.StatusBadRequest).JSON(contextMap)
	}

	result := db.Save(&task)

	if result.Error != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = result.Error
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}

	contextMap["task"] = task
	contextMap["message"] = "Task updated successfully"
	return ctx.Status(fiber.StatusOK).JSON(contextMap)
}

func DeleteTask(ctx *fiber.Ctx) error {
	contextMap := fiber.Map{
		"message":    "Task deleted",
		"statusText": "Ok",
	}
	db := database.DbConn
	taskId := ctx.Params("id")
	var task models.Task

	db.First(&task, taskId)

	if task.ID == 0 {
		contextMap["statusText"] = "Not Found"
		contextMap["message"] = "Task not found"
		return ctx.Status(fiber.StatusNotFound).JSON(contextMap)
	}

	result := db.Delete(&task)

	if result.Error != nil {
		contextMap["statusText"] = "Internal Server Error"
		contextMap["message"] = "Error deleting the task"
		return ctx.Status(fiber.StatusInternalServerError).JSON(contextMap)
	}

	contextMap["message"] = "Task deleted successfully"
	return ctx.Status(fiber.StatusOK).JSON(contextMap)
}
