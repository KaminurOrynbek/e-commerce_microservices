package main

import (
	"database/sql"
	"fmt"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/config"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/handler/http"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/repository/postgres"
	"github.com/gin-gonic/gin"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()

	// PostgreSQL connection
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	// repositories
	productRepo := postgres.NewProductRepository(db)
	categoryRepo := postgres.NewCategoryRepository(db)

	// handlers
	productHandler := http.NewProductHandler(productRepo)
	categoryHandler := http.NewCategoryHandler(categoryRepo)

	// Gin router
	router := gin.Default()

	// routes
	productHandler.RegisterRoutes(router)
	categoryHandler.RegisterRoutes(router)

	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
