package main

import (
	"os"
	"test-website/server/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip} ${id} [${time_rfc3339}] \"${method} ${uri}\" ${status} ${method} ${bytes_out}\n",
	}))
	e.Use(middleware.Recover())

	// Unauthenticated Routes
	e.Static("/", "dist")
	e.POST("/api/register", routes.RegisterUser)
	e.POST("/api/login", routes.LoginUser)

	// Authenticated routes
	r := e.Group("/api/restricted")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}))

	r.GET("/user/:id", routes.GetUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
