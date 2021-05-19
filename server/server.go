package main

import (
	"os"
	"test-website/server/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Echo instance
	e := echo.New()

	// Route logger
	logger := log.New("ROUTE")
	logger.SetHeader("[${time_rfc3339}] [${level}] ${short_file} L${line} ${message}")

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] \"${method} ${uri}\" ${status}\n",
	}))
	e.Use(middleware.Recover())

	// Serve Static UI Files (could be separated to another server)
	e.Static("/", "dist")

	// Authentication Routes
	e.POST("/api/register", func(c echo.Context) error {
		return routes.RegisterMember(c, logger)
	})
	e.POST("/api/login", func(c echo.Context) error {
		return routes.LoginMember(c, logger)
	})

	// Board Routes
	e.GET("/api/board", func(c echo.Context) error {
		return routes.GetBoards(c, logger)
	}, authenticated())
	e.GET("/api/board/:board_gid", func(c echo.Context) error {
		return routes.GetBoardById(c, logger)
	}, authenticated())
	e.PUT("/api/board/:board_gid", func(c echo.Context) error {
		return routes.EditBoard(c, logger)
	}, authenticated())
	e.POST("/api/board", func(c echo.Context) error {
		return routes.CreateBoard(c, logger)
	}, authenticated())
	e.DELETE("/api/board/:board_gid", func(c echo.Context) error {
		return routes.DeleteBoard(c, logger)
	}, authenticated())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func authenticated() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "header:" + echo.HeaderAuthorization,
	})
}
