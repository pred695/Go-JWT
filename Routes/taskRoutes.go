package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/Controllers"
	"github.com/pred695/Go-JWT/MiddleWare"
)

func SetUpTaskRoutes(app *fiber.App) {
	private := app.Group("/private")
	private.Use(MiddleWare.VerifyUser)
	private.Get("/tasks", Controllers.ListTasks)
	private.Post("/tasks", Controllers.CreateTask)
	private.Put("/tasks/:id", Controllers.UpdateTask)
	private.Delete("/tasks/:id", Controllers.DeleteTask)
}
