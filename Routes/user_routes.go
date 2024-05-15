package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/MiddleWare"
	"github.com/pred695/Go-JWT/controllers"
)

func SetUpUserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Post("/login", controllers.LoginUser)
	app.Post("/register", controllers.RegisterUser)

	private := app.Group("/private")
	private.Use(MiddleWare.VerifyUser)
	private.Get("/refresh", controllers.RefreshToken)
	private.Get("/logout", controllers.LogOutUser)
}
