package main

import "github.com/labstack/echo"

func main() {
	serverInstance := echo.New()
	GetRoutes(serverInstance)

	serverInstance.Logger.Fatal(serverInstance.Start(":8080"))
}
