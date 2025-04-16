package main

import (
	"fmt"
	"github.com/KaminurOrynbek/e-commerce_microservices/api-gateway/config"
	"github.com/KaminurOrynbek/e-commerce_microservices/api-gateway/internal/handler"
	"github.com/KaminurOrynbek/e-commerce_microservices/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.NewConfig()

	inventoryHandler := handler.NewInventoryHandler(cfg.Services.InventoryServiceURL)
	orderHandler := handler.NewOrderHandler(cfg.Services.OrderServiceURL)

	router := gin.Default()

	router.Use(middleware.Logger())
	router.Use(middleware.AuthMiddleware())

	api := router.Group("/api")
	{
		inventory := api.Group("/inventory/*path")
		{
			inventory.Any("", inventoryHandler.ProxyRequest)
		}

		orders := api.Group("/orders/*path")
		{
			orders.Any("", orderHandler.ProxyRequest)
		}
	}

	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("API Gateway starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
