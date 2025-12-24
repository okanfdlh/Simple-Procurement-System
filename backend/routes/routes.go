package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func Setup(app *fiber.App) {
	// Swagger (PUBLIC)
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	// Auth (PUBLIC)
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// Protected routes (JWT)
	protected := api.Group("/", middlewares.JWTProtected())

	// Items
	protected.Get("/items", controllers.GetItems)
	protected.Post("/items", controllers.CreateItem)
	protected.Put("/items/:id", controllers.UpdateItem)
	protected.Delete("/items/:id", controllers.DeleteItem)

	// Suppliers
	protected.Get("/suppliers", controllers.GetSuppliers)
	protected.Post("/suppliers", controllers.CreateSupplier)
	protected.Put("/suppliers/:id", controllers.UpdateSupplier)
	protected.Delete("/suppliers/:id", controllers.DeleteSupplier)

	// Purchasing
	protected.Post("/purchasing", controllers.CreatePurchasing)
}
