# Weather App - GOTTH Stack

A modern weather application built with Go + Templ + Tailwind CSS + HTMX using the Fiber framework.

## 🚀 Quick Start

### 1. Clone & Setup
```bash
git clone <your-repo>
cd weather-app-gotth
```

### 2. Install Dependencies
```bash
make install
```

### 3. Setup Environment
```bash
cp .env.example .env
nano .env  # Add your API key from weatherapi.com
```

### 4. Run Application
```bash
# With hot reload (recommended)
make dev

# Or normal mode
make run
```

### 5. Open Browser
Visit `http://localhost:8080`

## 📋 Prerequisites

- Go 1.21+
- Redis (optional - app works without it)
- Weather API key from [weatherapi.com](https://www.weatherapi.com/signup.aspx)

## 🏗️ Project Structure
```
weather-app-gotth/
├── cmd/server/           # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── logger/          # Logging setup
│   ├── cache/           # Redis cache
│   ├── weather/         # Weather domain logic
│   └── server/          # HTTP server & handlers
├── web/templates/       # Templ UI components
├── .env                 # Your secrets (gitignored)
├── .env.example         # Template
└── docker-compose.yml   # Docker setup
```

## 🔧 Environment Variables

Edit `.env` file:
```bash
SERVER_PORT=:8080
REDIS_ADDRESS=localhost:6379
API_KEY=your_api_key_here  # Get from weatherapi.com
```

## 🐳 Docker
```bash
# Start everything (app + Redis)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

## 📝 Make Commands
```bash
make install       # Install dependencies
make dev          # Development mode (hot reload)
make run          # Run application
make build        # Build binary
make clean        # Clean generated files
make test         # Run tests
make help         # Show all commands
```

## 🎨 Features

- ✅ Real-time weather data
- ✅ Beautiful responsive UI
- ✅ Redis caching
- ✅ Structured logging
- ✅ HTMX dynamic updates
- ✅ Docker support
- ✅ Environment-based config

## 📚 Tech Stack

- **Go** - Backend language
- **Fiber** - Web framework
- **Templ** - Type-safe templates
- **HTMX** - Dynamic interactions
- **Tailwind CSS** - Styling
- **Redis** - Caching
- **Logrus** - Structured logging
- **Viper** - Configuration management

## 🧪 Testing
```bash
# Run tests
make test

# Test specific package
go test ./internal/weather/...

# With coverage
go test -cover ./...
```

## 🐛 Troubleshooting

### API_KEY error
```bash
# Make sure you created .env file
cp .env.example .env
nano .env  # Add your API key
```

### Redis connection failed
```bash
# Start Redis
docker run -d -p 6379:6379 redis:7-alpine

# Or the app works without Redis (just logs a warning)
```

### Port already in use
```bash
# Change port in .env
SERVER_PORT=:9090
```

## 📖 Documentation

- [Fiber Docs](https://gofiber.io/)
- [Templ Guide](https://templ.guide/)
- [HTMX Docs](https://htmx.org/)
- [Weather API](https://www.weatherapi.com/docs/)

## 📄 License

MIT License

## 🙏 Credits

- Weather data from [WeatherAPI](https://www.weatherapi.com/)
- Built with the GOTTH stack
```

---

## 🎯 Complete Folder Structure
```
weather-app-gotth/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── logger/
│   │   └── logger.go
│   ├── cache/
│   │   └── redis.go
│   ├── weather/
│   │   ├── models.go
│   │   └── service.go
│   └── server/
│       ├── server.go
│       ├── routes.go
│       └── handlers.go
│
├── web/
│   └── templates/
│       ├── layout.templ
│       ├── index.templ
│       └── weather.templ
│
├── .env                    # Create this (gitignored)
├── .env.example            # Template provided
├── .gitignore             # Git ignore rules
├── go.mod                 # Dependencies
├── Makefile              # Build commands
├── Dockerfile            # Docker image
├── docker-compose.yml    # Docker orchestration
└── README.md             # Documentation
