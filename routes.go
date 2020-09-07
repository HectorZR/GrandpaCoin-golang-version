package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// GetRoutes set all the server routes
func GetRoutes(serverInstance *echo.Echo) {
	api := serverInstance.Group("/api")

	api.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]time.Time{
			"server_time": time.Now(),
		})
	})
}
