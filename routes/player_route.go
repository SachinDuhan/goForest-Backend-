package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/SachinDuhan/multiplayer/controllers"
)

func PlayerRoutes(app *fiber.App) {
	api := app.Group("/api/players")
	api.Post("/register", controllers.RegisterUser)
  api.Post("/login", controllers.Login)
  api.Get("/getPlayerInfo", controllers.GetPlayerInfo)
}
