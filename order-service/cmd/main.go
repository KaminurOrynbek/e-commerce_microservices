package main

import (
	"database/sql"
	"fmt"
	"github.com/KaminurOrynbek/e-commerce_microservices/order-service/internal/handler"
	"github.com/KaminurOrynbek/e-commerce_microservices/order-service/internal/repository"
	"github.com/KaminurOrynbek/e-commerce_microservices/order-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construct DB connection string from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbSSLMode)

	// Open a connection to the PostgreSQL database.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize the repository, use case, and handler layer for orders.
	orderRepo := repository.NewPgOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	// Create a Gin router and define routes.
	router := gin.Default()

	// Health check endpoint.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Order Service is running"})
	})

	// Order endpoints.
	ordersGroup := router.Group("/orders")
	{
		ordersGroup.POST("", orderHandler.CreateOrder)
		ordersGroup.GET("/:id", orderHandler.GetOrder)
		ordersGroup.PATCH("/:id", orderHandler.UpdateOrder)
		ordersGroup.GET("", orderHandler.ListOrdersByUser)
	}

	// Start the HTTP server on port 8083.
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
