package handlers

import (
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"

	"skport/go-api-server-echo/services"
)

type Handlers struct {
}

func NewHandler() *Handlers {
	h := new(Handlers)
	return h
}

func (h *Handlers) Init(e *echo.Echo, s *services.Services) {
	e.GET("/", s.Hello)
	e.GET("/albums", s.GetAlbums)
	e.GET("/albums/:id", s.GetAlbumByID)
	e.POST("/albums", s.PostAlbums)
}