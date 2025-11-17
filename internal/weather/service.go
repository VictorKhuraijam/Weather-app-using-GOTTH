package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"weather-app/internal/config"


	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// Service handles weather data fetching and caching
type Service struct {
	config config.Config
	redis  *redis.Client
	logger *logrus.Logger
}

// NewService creates a new weather service instance
func NewService(cfg *config.Config, redis *redis.Client, logger *logrus.Logger) *Service {
	return &Service{
		config: *cfg,
		redis:  redis,
		logger: logger,
	}
}

// GetWeather fetches weather data for a city with caching
// It first checks Redis cache, then falls back to API if needed
func (s *Service) GetWeather(ctx context.Context, city string) (*WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather:%s", city)

	// Try to get from cache first
	if s.redis != nil {
		cached, err := s.redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var weather WeatherResponse
			if err := json.Unmarshal([]byte(cached), &weather); err == nil {
				s.logger.WithFields(logrus.Fields{
					"city": city,
				}).Debug(" Cache hit")
				return &weather, nil
			}
		}
	}

	// Cache miss - fetch from API
	s.logger.WithFields(logrus.Fields{
		"city": city,
	}).Debug(" Fetching from API")

	weather, err := s.fetchFromAPI(city)
	if err != nil {
		return nil, err
	}

	// Cache the result
	s.cacheWeather(ctx, cacheKey, weather)

	return weather, nil
}

// fetchFromAPI fetches weather data from the external API
func (s *Service) fetchFromAPI(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://%s?key=%s&q=%s&aqi=no",
		s.config.APIURL, s.config.APIKey, city)

	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to fetch weather data: %v", errs[0])
	}

	if statusCode != 200 {
		return nil, fmt.Errorf("weather API returned status code: %d", statusCode)
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	return &weather, nil
}

// cacheWeather stores weather data in Redis cache
func (s *Service) cacheWeather(ctx context.Context, key string, weather *WeatherResponse) {
	if s.redis == nil {
		return
	}

	data, err := json.Marshal(weather)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Failed to marshal weather data for caching")
		return
	}

	expiry := time.Duration(s.config.RedisExpiryMin) * time.Minute
	if err := s.redis.Set(ctx, key, data, expiry).Err(); err != nil {
		s.logger.WithFields(logrus.Fields{
			"key":   key,
			"error": err.Error(),
		}).Warn("Failed to cache weather data")
		return
	}

	s.logger.WithFields(logrus.Fields{
		"key":    key,
		"expiry": expiry,
	}).Debug("Cached weather data")
}
