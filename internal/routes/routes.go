package routes

import (
	"github.com/gofiber/fiber/v2"
	"sqlc.dev/app/internal/handler"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/users", userHandler.Create)
	app.Get("/users/:id", userHandler.GetByID)
	app.Put("/users/:id", userHandler.Update)
	app.Delete("/users/:id", userHandler.Delete)
	app.Get("/users", userHandler.List)
}
