package main

import (
	log_d "log"
	"os"

	"github.com/null-bd/logger"
	"github.com/null-bd/microservice-name/config"
	"github.com/null-bd/microservice-name/config/database"
	"github.com/null-bd/microservice-name/config/router"
	"github.com/null-bd/microservice-name/internal/app"
)

func main() {
	// Get environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	// Load configuration
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log_d.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize log config
	log, err := config.GetLogger(cfg)
	if err != nil {
		log_d.Fatalf("Failed to initialize logger: %v", err)
	}

	// Initialize database connection
	db, err := database.NewPostgresConnection(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to initialize database", logger.Fields{
			"error": err,
		})
	}
	defer db.Close()

	// Initialize application
	application := app.NewApplication(log, cfg, db.Pool)

	// Initialize router with auth middleware
	router, err := router.NewRouter(log, cfg, &application.HealthHandler, &application.StaffHandler, &application.UserHandler)
	if err != nil {
		log.Fatal("Failed to initialize database", logger.Fields{
			"error": err,
		})
	}

	log.Info("Starting server", logger.Fields{"port": cfg.App.Port, "env": cfg.App.Env})
	if err := router.Run(); err != nil {
		log.Fatal("Failed to start server: ", logger.Fields{"error": err})
	}
}
