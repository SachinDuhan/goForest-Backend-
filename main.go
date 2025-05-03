package main

import (
	"github.com/SachinDuhan/multiplayer/routes"
  "github.com/SachinDuhan/multiplayer/config"
  "github.com/SachinDuhan/multiplayer/migrations"
  "github.com/SachinDuhan/multiplayer/middleware"
  "github.com/gofiber/fiber/v2"
)

func main() {
  // Initializing fiber app
	app := fiber.New()

  app.Use(middleware.ErrorHandler())

  // Connectig to database
  config.ConnectDatabase()
  // Initializing/migrating database
  migrations.InitDB()

	// Simple test routes
  routes.PlayerRoutes(app)
	// Start server
	app.Listen(":3000")
}

