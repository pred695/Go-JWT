package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/controllers"
	"github.com/pred695/Go-JWT/middleware"
)

func SetUpTaskroutes(app *fiber.App) {
	private := app.Group("/private")
	private.Use(middleware.VerifyUser)
	private.Get("/tasks", controllers.ListTasks)
	private.Post("/tasks", controllers.CreateTask)
	private.Put("/tasks/:id", controllers.UpdateTask)
	private.Delete("/tasks/:id", controllers.DeleteTask)
}
