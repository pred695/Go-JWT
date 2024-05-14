package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pred695/Go-JWT/Controllers"
	"github.com/pred695/Go-JWT/MiddleWare"
)

func SetUpUserRoutes(app *fiber.App) {
	app.Get("/users", Controllers.GetUsers)
	app.Post("/login", Controllers.LoginUser)
	app.Post("/register", Controllers.RegisterUser)

	private := app.Group("/private")
	private.Use(MiddleWare.VerifyUser)
	private.Get("/refresh", Controllers.RefreshToken)
	private.Get("/logout", Controllers.LogOutUser)
}
