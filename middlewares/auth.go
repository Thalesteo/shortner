package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/Thalesteo/trypgx/handlers"
	"github.com/Thalesteo/trypgx/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func SetupAuth() {
	store = session.New(session.Config{
		Storage: postgres.New(postgres.Config{
			ConnectionURI: os.Getenv("DB_SERVER_URL"),
			// Host:          "127.0.0.1",
			// Port:          5432,
			// Username:      "postgres",
			// Password:      "postgres",
			// Database:      "admin",
			Table:      "session",
			SSLMode:    "disable",
			Reset:      false,
			GCInterval: 10 * time.Second,
		}),
		CookieHTTPOnly: true,
		CookieSecure:   false,
		Expiration:     time.Hour * 5,
	})
}

func Auth() fiber.Handler {
	return AuthHandler
}

func AuthHandler(c *fiber.Ctx) error {
	// try to get session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Not Authorized")
	}

	// if session does not exist
	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Not Authorized")
	}

	// if session exist
	fmt.Println("passed by Auth")
	return c.Next()
}

func Login(c *fiber.Ctx) error {
	var data user

	// parse data from client
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something went wrong!")
	}

	// check if email exists in db
	var user models.User
	if !handlers.CheckUserEmail(data.Email, &user) {
		return c.Status(fiber.StatusUnauthorized).SendString("You don't have an account")
	}

	// check if client password is igual to db user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Password do not match")
	}

	// initiate a session
	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something went wrong!")
	}

	// setup session vars
	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.ID)

	// save session
	sessErr = sess.Save()
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something went wrong!")
	}

	// if everything went ok login
	return c.Status(fiber.StatusOK).SendString("Logged in")
}

func Logout(c *fiber.Ctx) error {
	// check if session exists if not than aready logged out
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).SendString("Logged out - no session")
	}

	// if session exists destroy session
	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Something went wrong!")
	}

	// user logged out
	return c.Status(fiber.StatusOK).SendString("Logged out")
}

func AuthHealthCheck(c *fiber.Ctx) error {
	// check if session exists
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Not Authorized")
	}

	// check is session key still valid
	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Not Authorized")
	}

	// if key is valid user aready logged in
	return c.Status(fiber.StatusOK).SendString("Authenticated")
}
