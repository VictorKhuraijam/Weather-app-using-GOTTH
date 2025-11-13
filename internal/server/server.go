package server

import (
	"weather-app/internal/config"
	"weather-app/internal/weather"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

// Server represents the HTTP server
type Server struct {
	app            *fiber.App
	config         config.Config
	weatherService *weather.Service
	logger         *logrus.Logger
}

// New creates a new server instance with all middleware and routes configured
func New(cfg config.Config, weatherService *weather.Service, log *logrus.Logger) *Server {
	// Create Fiber app with custom error handler
	app := fiber.New(fiber.Config{
		AppName: "Weather App GOTTH",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.WithFields(logrus.Fields{
				"error": err.Error(),
				"path":  c.Path(),
				"ip":    c.IP(),
			}).Error(" Request error")
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	srv := &Server{
		app:            app,
		config:         cfg,
		weatherService: weatherService,
		logger:         log,
	}

	// Setup middleware
	srv.setupMiddleware()

	// Setup routes
	srv.setupRoutes()

	return srv
}

// setupMiddleware configures all middleware
func (s *Server) setupMiddleware() {
	// Recover from panics
	s.app.Use(recover.New())

	// HTTP request logging
	s.app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
		Output: colorable.NewColorableStdout(),
	}))
}

// Start starts the HTTP server
func (s *Server) Start() {
	port := s.config.ServerPort
	if port == "" {
		port = ":8080"
	}

	s.logger.WithFields(logrus.Fields{
		"port": port,
	}).Info(" Server ready and listening")

	if err := s.app.Listen(port); err != nil {
		s.logger.Fatal(" Server failed to start:", err)
	}
}
