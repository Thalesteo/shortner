package main

import (
	"log"

	"github.com/Thalesteo/trypgx/db"
	"github.com/Thalesteo/trypgx/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Migrate()
}

func main() {
	app := fiber.New()

	// middlewares.SetupAuth()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user/:id", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)

	/*
		Client endpoints
		GET		/				template/index
		GET		/login			template/login
		POST  	/login 			Login
		GET		/register		template/register
		POST 	/register 		createUser
		GET		/recover		template/recover
		POST	/recover		RecoverPassword
		GET		/not-found		template/404
	*/

	/*
		Auth endpoints ALL middleware(auth)
		GET		/logout			Logout
	*/

	/*
		users endpoints ALL middleware(auth)
		GET  	/user/:id 		GetUser
		PUT  	/user/:id 		EditUser
		DELETE	/user/:id		DeleteUser
	*/

	/*
		link groups endpoints ALL middleware(auth)
		GET		/links			GetGroups
		POST  	/links			CreateGroup
		GET 	/links/:id 		GetGroup
		PUT  	/links/:id 		EditGroup
		DELETE	/links/:id		DeleteGroup
		DELETE	/Links			DeleteGroups
	*/

	/*
		link endpoints ALL middleware(auth)
		GET		/links/:id/:id			GetLink
		POST  	/links/:id/link			CreateLink
		PUT		/links/:id/:id			UpdateLink
		DELETE	/links/:id/:id			DeleteLink
		DELETE	/Links/:id				DeleteLinks
	*/

	log.Fatal(app.Listen(":3000"))
}
