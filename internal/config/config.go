package config

import (
	// "fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	// "github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	ServerPort        string `mapstructure:"SERVER_PORT"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword     string `mapstructure:"REDIS_PASS"`
	RedisDB           int    `mapstructure:"REDIS_DB"`
	RedisExpiryMin    int    `mapstructure:"REDIS_EXPIRY_MIN"`
	ContextTimeoutSec int    `mapstructure:"CONTEXT_TIMEOUT_SEC"`
	APIKey            string `mapstructure:"API_KEY"`
	APIURL            string `mapstructure:"API_URL"`
}

func Load() *Config{
	godotenv.Load()

	return &Config{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		RedisAddress: getEnv("REDIS_ADDRESS", ""),
		RedisPassword: getEnv("REDIS_PASS",""),
		RedisDB: getEnv("REDIS_DB", 0),
		RedisExpiryMin: getEnv("REDIS_EXPIRY_MIN", 30),
		ContextTimeoutSec: getEnv("CONTEXT_TIMEOUT_SEC", 10),
		APIKey: getEnv("API_KEY","" ),
		APIURL: getEnv("API_URL", ""),
	}
}

func getEnv[T any](key string, defaultValue T) T {
	 value := os.Getenv(key);
	if value == "" {
		return defaultValue
	}
	var result any

	switch any(defaultValue).(type){
	case string:
		result = value

	case int:
		v, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		result = v

	default:
		//Unsupported type -> fallback to default
		return defaultValue
	}
	return result.(T)
}








// Load reads configuration from .env and environment variables
// Environment variables take precedence over .env file
// Example: SERVER_PORT in environment overrides SERVER_PORT in .env
// func Load() Config {
// 	//Load .env file if it exists (optional)
// 	//This will not override exsiting environment variables
// 	if err := godotenv.Load(); err != nil {
// 		fmt.Println("No .env file found, using environment variables")
// 	}

// 	//Configure Viper to read from environment variables
// 	viper.AutomaticEnv()

// 	// Set default values
// 	viper.SetDefault("SERVER_PORT", ":8080")
// 	viper.SetDefault("REDIS_ADDRESS", "localhost:6379")
// 	viper.SetDefault("REDIS_PASS", "")
// 	viper.SetDefault("REDIS_DB", 0)
// 	viper.SetDefault("REDIS_EXPIRY_MIN", 30)
// 	viper.SetDefault("CONTEXT_TIMEOUT_SEC", 10)
// 	viper.SetDefault("API_URL", "api.weatherapi.com/v1/current.json")

// 	viper.BindEnv("SERVER_PORT")
//     viper.BindEnv("REDIS_ADDRESS")
//     viper.BindEnv("REDIS_PASS")
//     viper.BindEnv("REDIS_DB")
//     viper.BindEnv("REDIS_EXPIRY_MIN")
//     viper.BindEnv("CONTEXT_TIMEOUT_SEC")
// 	viper.BindEnv("API_KEY")

// 	// Allow environment variables to override config file
// 	// viper.SetEnvPrefix("APP")
// 	// viper.AutomaticEnv()

// 	var config Config

// 	// if err := viper.ReadInConfig(); err != nil {
// 	// 	// Config file is optional - can run with env vars only
// 	// 	fmt.Printf(" Config file not found, using environment variables: %v\n", err)
// 	// }

// 	// Unmarshal environment variables into Config struct
// 	if err := viper.Unmarshal(&config); err != nil {
// 		fmt.Printf(" Failed to parse config: %v\n", err)
// 		panic(err)
// 	}

// 	// Validate required fields
// 	if config.APIKey == "" {
// 		fmt.Println(" ERROR: API_KEY is required!")
// 		fmt.Println(" Get your free API key from: https://www.weatherapi.com/signup.aspx")
// 		fmt.Println(" export API_KEY=your_actual_key")
// 		os.Exit(1)
// 	}

// 	return config
// }
