package presentation

import (
	"github.com/davidyanez27/go-gorm-restapi/src/presentation/auth"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	authRoutes := auth.SetupAuthRoutes()
	app.Mount("/product", authRoutes)
}
