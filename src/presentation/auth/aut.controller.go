package auth

import (
	"fmt"

	"github.com/davidyanez27/go-gorm-restapi/src/data/models"
	"github.com/davidyanez27/go-gorm-restapi/src/db"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        string
	FirstName string
	LastName  string
}

func handleCreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	if err := db.DB.Create(&product).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println(product)

	return c.Status(fiber.StatusOK).JSON(product)
}

func handleGetProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := db.DB.Find(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println(product)

	return c.Status(fiber.StatusOK).JSON(product)
}

func handleUpdateProduct(c *fiber.Ctx) error {
	// Get the product ID from the URL parameters
	id := c.Params("id")

	// Ensure the product exists before attempting to update
	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	// Parse the incoming request body into the existingProduct struct
	if err := c.BodyParser(&existingProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Save the updated product
	if err := db.DB.Save(&existingProduct).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(existingProduct)
}

func handleDeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Ensure the product exists before attempting to update
	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, product.ID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	// Update the product
	if err := db.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
