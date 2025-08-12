package http

import (
	"codelabs-backend-fiber/internal/user/domain"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, uc domain.UserUsecase) {
	handler := NewUserHandler(uc)

	api := app.Group("/api")
	users := api.Group("/users")

	// User routes
    users.Get("/", handler.GetAll)
    users.Get("/:id", handler.GetByID)
    users.Post("/", handler.Create)
    
    // Auth routes
    api.Post("/register", handler.Create)
	api.Post("/login", handler.Login)
}
