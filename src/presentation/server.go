package presentation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	App    *fiber.App
	Port   int
	Routes func(app *fiber.App)
}

func Server(port int, routes func(app *fiber.App)) *server {
	return &server{
		App:    fiber.New(),
		Port:   port,
		Routes: routes,
	}
}

func (s *server) Start() {

	// s.Routes(s.App)
	s.Routes(s.App)

	// Middleware for parsing JSON
	s.App.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})

	// Middleware for parsing URL-encoded form data
	s.App.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/x-www-form-urlencoded")
		return c.Next()
	})

	s.App.Listen(fmt.Sprintf(":%d", s.Port))

}
