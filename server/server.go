package main

import (
	"organizr/server/routes"
	"os"

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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
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

	// Task Column Routes
	e.GET("/api/column/:board_gid", func(c echo.Context) error {
		return routes.GetColumns(c, logger)
	}, authenticated())
	e.GET("/api/column/:board_gid/:column_gid", func(c echo.Context) error {
		return routes.GetColumnById(c, logger)
	}, authenticated())
	e.PUT("/api/column/:board_gid/:column_gid", func(c echo.Context) error {
		return routes.EditColumn(c, logger)
	}, authenticated())
	e.POST("/api/column/:board_gid", func(c echo.Context) error {
		return routes.CreateColumn(c, logger)
	}, authenticated())
	e.DELETE("/api/column/:board_gid/:column_gid", func(c echo.Context) error {
		return routes.DeleteColumn(c, logger)
	}, authenticated())

	// Task Routes
	e.GET("/api/task/:column_id", func(c echo.Context) error {
		return routes.GetTasks(c, logger)
	}, authenticated())
	e.GET("/api/task/:column_id/:task_id", func(c echo.Context) error {
		return routes.GetTaskById(c, logger)
	}, authenticated())
	e.PUT("/api/task/:column_id/:task_id", func(c echo.Context) error {
		return routes.EditTask(c, logger)
	}, authenticated())
	e.POST("/api/task/:column_id", func(c echo.Context) error {
		return routes.CreateTask(c, logger)
	}, authenticated())
	e.DELETE("/api/task/:column_id/:task_id", func(c echo.Context) error {
		return routes.DeleteTask(c, logger)
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
