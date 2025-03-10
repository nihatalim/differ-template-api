package main

import (
	"log"

	"differ-template-api/internal/health"
	"differ-template-api/pkg/config"
	"differ-template-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg := config.New()

	// Initialize Fiber app
	app := fiber.New()

	// Add middleware
	app.Use(middleware.Logger())

	// Initialize services and handlers
	healthService := health.NewService()
	healthHandler := health.NewHandler(healthService)

	// Routes
	app.Get("/", healthHandler.HandleHealthCheck)

	// Start server
	log.Fatal(app.Listen(cfg.Server.Port))
} 