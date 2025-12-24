package main

import (
	"log"

	"backend/config"
	"backend/models"
	"backend/routes"
	"backend/seeders"

	_ "backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// @title Simple Procurement API
// @version 1.0
// @description API for Simple Procurement System
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// Connect database
	config.ConnectDB()

	// Auto migrate tables
	config.DB.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Item{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	)
	seeders.RunAll()

	// Start app
	app := fiber.New()
	// ✅ CORS CONFIG
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // sementara (DEV)
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
