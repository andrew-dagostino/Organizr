package main

import (
	"organizr/server/routes"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		AllowOrigins: []string{"*"},
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

	// Authenticated Paths
	r := e.Group("/api/r")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	// Board Routes
	r.GET("/board", func(c echo.Context) error { return routes.GetBoards(c, logger) })
	r.GET("/board/:board_gid", func(c echo.Context) error { return routes.GetBoardById(c, logger) })
	r.PUT("/board/:board_gid", func(c echo.Context) error { return routes.EditBoard(c, logger) })
	r.POST("/board", func(c echo.Context) error { return routes.CreateBoard(c, logger) })
	r.DELETE("/board/:board_gid", func(c echo.Context) error { return routes.DeleteBoard(c, logger) })

	// Task Column Routes
	r.GET("/column", func(c echo.Context) error { return routes.GetColumns(c, logger) })
	r.GET("/column/:column_gid", func(c echo.Context) error { return routes.GetColumnById(c, logger) })
	r.PUT("/column/:column_gid", func(c echo.Context) error { return routes.EditColumn(c, logger) })
	r.POST("/column", func(c echo.Context) error { return routes.CreateColumn(c, logger) })
	r.DELETE("/column/:column_gid", func(c echo.Context) error { return routes.DeleteColumn(c, logger) })

	// Task Routes
	r.GET("/task", func(c echo.Context) error { return routes.GetTasks(c, logger) })
	r.GET("/task/:task_gid", func(c echo.Context) error { return routes.GetTaskById(c, logger) })
	r.PUT("/task/:task_gid", func(c echo.Context) error { return routes.EditTask(c, logger) })
	r.POST("/task", func(c echo.Context) error { return routes.CreateTask(c, logger) })
	r.DELETE("/task/:task_gid", func(c echo.Context) error { return routes.DeleteTask(c, logger) })

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
