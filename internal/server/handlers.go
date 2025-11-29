package server

import (
	"context"
	"fmt"
	"time"

	"weather-app/web/templates"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// handleIndex renders the home page
func (s *Server) handleIndex(c *fiber.Ctx) error {
	s.logger.Debug("Rendering index page")
	component := templates.IndexPage()
	return s.render(c, component)
}

// handleWeather fetches and displays weather data
func (s *Server) handleWeather(c *fiber.Ctx) error {
	city := c.Query("city")
	if city == "" {
		s.logger.Warn("Weather request missing city parameter")
		// return c.Status(fiber.StatusBadRequest). //4xx response
		// 	SendString("<div class='p-4 bg-red-100 text-red-700 rounded'>Please provide a city name</div>")
		//HTMX DOES NOT SWAP anything when the response status is not 2xx unless configured to handle them.
		return c.SendString("<div class='p-4 bg-red-100 text-red-700 rounded'>Please provide a city name</div>")
	}

	// Log the request with context
	s.logger.WithFields(logrus.Fields{
		"city": city,
		"ip":   c.IP(),
	}).Info(" Weather request")

	// Fetch weather data
	weather, err := s.weatherService.GetWeather(c.Context(), city)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"city":  city,
			"error": err.Error(),
		}).Error(" Failed to fetch weather")
			
		// return c.Status(fiber.StatusInternalServerError).
		// 	SendString(fmt.Sprintf("<div class='p-4 bg-red-100 text-red-700 rounded'>Error: %v</div>", err))
			return c.SendString(fmt.Sprintf("<div class='p-4 bg-red-100 text-red-700 rounded'>Error: %v</div>", err))
	}
	fmt.Println("Weather is :",weather)

	// Log successful fetch
	s.logger.WithFields(logrus.Fields{
		"city": weather.Location.Name,
		"temp": weather.Current.TempC,
	}).Info(" Weather data retrieved")

	// Render weather card component
	component := templates.WeatherCard(weather)
	return s.render(c, component)
}

// handleHealth returns application health status
func (s *Server) handleHealth(c *fiber.Ctx) error {
	redisStatus := "disconnected"

	// Check Redis connection if available
	if s.weatherService != nil {
		_, cancel := context.WithTimeout(c.Context(), 2*time.Second)
		defer cancel()

		// Try a simple operation to verify Redis is working
		// Note: This is a simplified check, in production you'd access redis through the service
		redisStatus = "unknown"
	}

	status := "healthy"
	currentTime := time.Now().Format(time.RFC3339)
	version := "1.0.0"

	// return c.JSON(fiber.Map{
	// 	"status":    "healthy",
	// 	"redis":     redisStatus,
	// 	"timestamp": time.Now(),
	// 	"version":   "1.0.0",
	// })
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.SendString(fmt.Sprintf(
		`<div style="
				padding: 1rem;
				display: flex;
				align-items: center;
				flex-direction: column;
				background-color: #d1fae5;
				color: #047857;
				border-radius: 0.25rem;
				">
			<p>Status: %s</p>
			<p>Redis: %s</p>
			<p>Timestamp: %s</p>
			<p>Version: %s</p>
		</div>`,
		status,
		redisStatus,
		currentTime,
		version,
	))

}

// render is a helper function to render Templ components
func (s *Server) render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
