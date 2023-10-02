package main

import (
	"github.com/Thalesteo/shortner/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.GetUser)
}
