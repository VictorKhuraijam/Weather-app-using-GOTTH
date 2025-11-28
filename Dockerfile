# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ files and build
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o weather-app .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/weather-app .

# Copy .env.example as reference (users should provide their own .env)
COPY --from=builder /app/.env.example .

EXPOSE 8080

# Note: Users should provide environment variables via docker-compose or -e flags
# Example: docker run -e API_KEY=your_key weather-app
CMD ["./weather-app"]
