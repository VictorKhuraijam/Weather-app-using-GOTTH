package main

import (
	"weather-app/internal/cache"
	"weather-app/internal/config"
	"weather-app/internal/logger"
	"weather-app/internal/server"
	"weather-app/internal/weather"
)

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
