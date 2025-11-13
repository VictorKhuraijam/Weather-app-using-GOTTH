package cache

import (
	"context"
	"time"
	"weather-app/internal/config"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// InitRedis initializes Redis connection with health check
// Returns nil if Redis is unavailable - app will work without cache
// This allows graceful degradation instead of crashing
func InitRedis(config config.Config, logger *logrus.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	// Test connection with timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.ContextTimeoutSec)*time.Second,
	)
	defer cancel()

	// Ping Redis to verify connection
	if err := client.Ping(ctx).Err(); err != nil {
		logger.WithFields(logrus.Fields{
			"address": config.RedisAddress,
			"error":   err.Error(),
		}).Warn(" Redis connection failed - running without cache")
		return nil
	}

	logger.WithFields(logrus.Fields{
		"address": config.RedisAddress,
		"db":      config.RedisDB,
	}).Info(" Redis connected")

	return client
}
