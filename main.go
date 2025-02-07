package main

import (
	"fmt"
	"log"
	"crud/config"
	"crud/middleware"
	"crud/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect Database
	config.ConnectDatabase()

	// Create Fiber App
	app := fiber.New()

	// Middleware
	app.Use(middleware.Logger())

	// Routes
	routes.UserRoutes(app)

	// Start Server
	fmt.Println("Server is running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
