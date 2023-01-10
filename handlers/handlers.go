package handlers

import (
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"

	"github.com/swaggo/echo-swagger"

	"skport/go-api-server-echo/services"
)

type Handlers struct {
}

func NewHandler() *Handlers {
	h := new(Handlers)
	return h
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func (h *Handlers) Init(e *echo.Echo, s *services.Services) {
	e.GET("/healthcheck", h.HealthCheck)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", s.Hello)
	e.GET("/albums", s.GetAlbums)
	e.GET("/albums/:id", s.GetAlbumByID)
	e.POST("/albums", s.PostAlbums)
}