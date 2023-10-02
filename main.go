package main

import (
	"os"

	"github.com/Thalesteo/shortner/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// setup
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// config
	app.Static("/", "./public")

	// routes
	Routes(app)

	// start
	app.Listen(":" + os.Getenv("PORT"))
}
