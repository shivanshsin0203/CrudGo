package main

import (
	"os"

	"analayticbackend/config"
	"analayticbackend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {

	 // init db
	err := config.InitDB()
	if err != nil {
		return err
	}

	// defer closing db
	defer config.CloseDB()

	// create app
	app := fiber.New()

	// add basic middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// add routes
	router.AddBookGroup(app)

	// start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3001"
	}
	app.Listen(":" + port)

	return nil
}