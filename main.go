package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"skport/go-api-server-echo/handlers"
	"skport/go-api-server-echo/services"
	//"skport/go-api-server-echo/configs"
	//"skport/go-api-server-echo/repository"
	//"skport/go-api-server-echo/stores"
)

func main() {
	log.Println("Starting.")

	// Configs
	//c := configs.NewConfigs()

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

	// Services
	s := services.NewServices()

	// Handlers (Routing)
	h := handlers.NewHandler()
	h.Init(e, s)

	// Stores
	//s := stores.NewStore()

	// Repository
	//r := repository.NewRepository()

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
