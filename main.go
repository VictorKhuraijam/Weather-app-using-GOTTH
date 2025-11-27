package main

import (
	"weather-app/internal/cache"
	"weather-app/internal/config"
	"weather-app/internal/logger"
	"weather-app/internal/server"
	"weather-app/internal/weather"
)

// Dependency Injection means:
// ✔ A struct does NOT create the things it depends on
// ✔ main.go provides the dependencies
// ✔ service.go only uses the dependencies
// ✔ cleaner code, easier testing, better architecture
// ✔ predictable memory layout and resource control
// ✔ enables replacing Redis, logger, or config easily

func main() {
	// Initialize configuration
	cfg := config.Load()

	// Initialize logger
	log := logger.InitLogger()
	log.Info("Starting Weather App (GOTTH Stack)")

	// Initialize Redis cache
	redisClient := cache.InitRedis(cfg, log)

	// Initialize weather service
	weatherService := weather.NewService(cfg, redisClient, log)

	// Initialize and start server
	srv := server.New(cfg, weatherService, log)
	srv.Start()
}
