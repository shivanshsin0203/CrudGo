package routes

import (
	"github.com/gofiber/fiber/v2"
	"crud/handler"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", handlers.GetUsers)
	api.Post("/users", handlers.CreateUser)
	api.Get("/users/:id", handlers.GetUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)
}
