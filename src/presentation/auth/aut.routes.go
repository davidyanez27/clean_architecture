package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes() *fiber.App {
	router := fiber.New()

	// router.Get("/login/:id", handleUser)
	router.Post("/create", handleCreateProduct)
	router.Get("/find", handleGetProduct)
	router.Put("/update", handleUpdateProduct)
	router.Delete("/delete", handleDeleteProduct)

	return router
}
