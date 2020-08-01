package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"

	"test-website/server/types"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Static("/", "dist")
	e.POST("/api/register", createUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func createUser(c echo.Context) (err error) {
	username := strings.TrimSpace(c.FormValue("username"))
	password := strings.TrimSpace(c.FormValue("password"))

	if len(username) < 1 || len(username) > 32 {
		return c.JSON(http.StatusBadRequest, &types.Error{
			Code:  "create_failed",
			Error: "Failed to create user",
		})
	}

	hash, err := hashPassword(password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &types.Error{
			Code:  "create_failed",
			Error: "Failed to create user",
		})
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return c.JSON(http.StatusInternalServerError, &types.Error{
			Code:  "create_failed",
			Error: "Failed to create user",
		})
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &types.Error{
			Code:  "create_failed",
			Error: "Failed to create user",
		})
	}
	defer tx.Rollback(context.Background())

	r_name := ""
	r_id := -1

	err = tx.QueryRow(context.Background(),
		`INSERT INTO site_user (username, password) VALUES ($1, $2)
			RETURNING id, username;`,
		username, hash,
	).Scan(&r_id, &r_name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &types.Error{
			Code:  "create_failed",
			Error: "Failed to create user",
		})
	}

	u := &types.User{
		Id:       r_id,
		Username: r_name,
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, &types.Error{
			Code:  "create_failed",
			Error: "Failed to create user",
		})
	}

	return c.JSON(http.StatusOK, u)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
