package main

import (
	"github.com/zaidalmaghfur/go-app/config"
	route "github.com/zaidalmaghfur/go-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()

	route.Init(app)

	app.Listen(":8080")
}
