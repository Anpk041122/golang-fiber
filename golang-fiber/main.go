package main

import (
	"log"

	"golang-fiber/config"
	"golang-fiber/setRoutes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(cors.New())

	setRoutes.RoutesUser(app)
	setRoutes.RoutesTodo(app)
	log.Fatal(app.Listen(":3000"))
}
