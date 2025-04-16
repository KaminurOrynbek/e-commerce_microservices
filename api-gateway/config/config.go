package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server   *ServerConfig
	Services *ServicesConfig
}

type ServerConfig struct {
	Port string
}

type ServicesConfig struct {
	InventoryServiceURL string
	OrderServiceURL     string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	return &Config{
		Server: &ServerConfig{
			Port: getEnv("SERVER_PORT", "8000"),
		},
		Services: &ServicesConfig{
			InventoryServiceURL: getEnv("INVENTORY_SERVICE_URL", "http://localhost:8080"),
			OrderServiceURL:     getEnv("ORDER_SERVICE_URL", "http://localhost:8081"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
