package handlers

import (
	"context"
	"crud/config"
	"crud/models"
	"github.com/gofiber/fiber/v2"
)

// Get all users
func GetUsers(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get users"})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error scanning users"})
		}
		users = append(users, user)
	}
	return c.JSON(users)
}

// Create user
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := config.DB.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
}

// Get user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	err := config.DB.QueryRow(context.Background(), "SELECT id, name, email FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// Update user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := config.DB.Exec(context.Background(), "UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated"})
}

// Delete user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.DB.Exec(context.Background(), "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted"})
}
