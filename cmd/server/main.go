package main

import (
	"codelabs-backend-fiber/config"
	"codelabs-backend-fiber/infrastructure/postgres"
	"codelabs-backend-fiber/internal/user/delivery/http"
	"codelabs-backend-fiber/internal/user/repository"
	"codelabs-backend-fiber/internal/user/usecase"
	database "codelabs-backend-fiber/migrations"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()
	cfg := config.AppConfig

	// Create DSN string from config
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	db := postgres.NewGormDB(dsn)

	// AutoMigrate Database
	database.Migrate(db)

	// Init Module
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	app := fiber.New()

	// Register Routes
	http.RegisterRoutes(app, userUsecase)

	// Started the apps
	app.Listen(":" + cfg.Port)
}
