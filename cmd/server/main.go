package main

import (
	"log"

	"github.com/yudin-artem/avito_test_2025/internal/config"
	"github.com/yudin-artem/avito_test_2025/internal/handler"
	"github.com/yudin-artem/avito_test_2025/internal/repository"
	"github.com/yudin-artem/avito_test_2025/internal/service"
)

func main() {
	db, err := repository.DBInit()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)	
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handler := handler.NewHandler(services)

	router := handler.InitRoutes()

	port, err := config.Get("SERVER_PORT")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}