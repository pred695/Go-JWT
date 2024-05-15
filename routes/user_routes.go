package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/controllers"
	"github.com/pred695/Go-JWT/middleware"
)

func SetUpUserroutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Post("/login", controllers.LoginUser)
	app.Post("/register", controllers.RegisterUser)

	private := app.Group("/private")
	private.Use(middleware.VerifyUser)
	private.Get("/refresh", controllers.RefreshToken)
	private.Get("/logout", controllers.LogOutUser)
}
