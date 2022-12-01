package setRoutes

import (
	"golang-fiber/todo"

	"github.com/gofiber/fiber/v2"
)

func RoutesTodo(app *fiber.App) {
	app.Get("/todo/:userid", todo.GetAll)

	app.Post("/todo", todo.Post)

	app.Put("/todo/:id", todo.UpdateContent)
	app.Put("/todo/:id", todo.UpdateDone)

	app.Delete("/todo/:id", todo.Delete)
}
