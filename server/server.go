package main

import (
	"test-website/server/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Static("/", "dist")
	e.POST("/api/register", routes.CreateUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
