package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"sqlc.dev/app/config"
	"sqlc.dev/app/internal/handler"
	"sqlc.dev/app/internal/logger"
	"sqlc.dev/app/internal/middleware"
	"sqlc.dev/app/internal/models"
	"sqlc.dev/app/internal/repository"
	"sqlc.dev/app/internal/routes"
	"sqlc.dev/app/internal/service"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	log := logger.Get()
	defer log.Sync()

	log.Info("Starting server setup...")

	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	dbConn, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to open database connection", zap.Error(err))
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		log.Fatal("Failed to ping database", zap.Error(err))
	}
	log.Info("Database connection established successfully")

	// Run simple migration
	log.Info("Running migrations...")
	migrationQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		dob DATE NOT NULL
	);`
	_, err = dbConn.Exec(migrationQuery)
	if err != nil {
		log.Fatal("Failed to run database migrations", zap.Error(err))
	}
	log.Info("Migrations completed successfully")

	// Initialize Validator
	models.InitValidator()

	// Initialize layers
	userRepo := repository.NewUserRepository(dbConn)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	// Setup GoFiber App
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Error("Unhandled error occurred", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		},
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	// Routes
	routes.SetupRoutes(app, userHandler)

	// Fallback route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not Found",
		})
	})

	// Start server
	log.Info(fmt.Sprintf("Server is starting on port %s", cfg.ServerPort))
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatal("Server failed to start", zap.Error(err))
	}
}
type dummy struct{} // avoid empty main errors if compiled as package
