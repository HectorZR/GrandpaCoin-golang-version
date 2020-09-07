package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetRoutes set all the server routes
func GetRoutes(serverInstance *echo.Echo) {
	serverInstance.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world!")
	})
}
