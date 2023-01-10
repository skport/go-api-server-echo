package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// HealthCheck godoc
// @Summary      Show HealthCheck
// @Description  Show HealthCheck
// @Tags         System
// @Accept       */*
// @Produce      plain
// @Success      200 {string} string "OK"
// @Router       /healthcheck [get]
func (h *Handlers) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}