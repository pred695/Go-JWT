package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/MiddleWare"
	"github.com/pred695/Go-JWT/controllers"
)

func SetUpTaskRoutes(app *fiber.App) {
	private := app.Group("/private")
	private.Use(MiddleWare.VerifyUser)
	private.Get("/tasks", controllers.ListTasks)
	private.Post("/tasks", controllers.CreateTask)
	private.Put("/tasks/:id", controllers.UpdateTask)
	private.Delete("/tasks/:id", controllers.DeleteTask)
}
