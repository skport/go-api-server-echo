package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"

	"skport/go-api-server-echo/handlers"
	"skport/go-api-server-echo/repository"
	"skport/go-api-server-echo/services"
)

func main() {
	log.Println("Starting.")

	// Environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// ------
	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	// ------

	// Switch data store using Repository pattern
	env := os.Getenv("APP_ENV")
	var rp repository.Repository
	if env == "development" {
		log.Println("Selected data store : MySQL")
		rp = repository.NewMySQLRepository()
	} else {
		log.Println("Selected data store : InMemory")
		rp = repository.NewInMemoryRepository()
	}

	// Services
	s := services.NewServices(&rp)

	// Handlers (Routing)
	h := handlers.NewHandler()
	h.Init(e, s)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
