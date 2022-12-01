package setRoutes

import (
	"golang-fiber/user"

	"github.com/gofiber/fiber/v2"
)

func RoutesUser(app *fiber.App) {

	app.Get("/user/:username", user.GetUserName)
	app.Get("/user/:username/:password", user.GetPassword)

	app.Post("/user", user.Post)

	app.Put("/user/:id", user.Put)

	app.Delete("/user/:id", user.Delete)
}
