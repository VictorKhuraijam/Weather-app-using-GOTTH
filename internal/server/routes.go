package server

// setupRoutes registers all HTTP routes
func (s *Server) setupRoutes() {
	// Home page
	s.app.Get("/", s.handleIndex)

	// Weather endpoint (HTMX target)
	s.app.Get("/weather", s.handleWeather)

	// Health check endpoint
	s.app.Get("/health", s.handleHealth)
}
