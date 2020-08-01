package routes

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Registers a new user using the POST's username and password
func RegisterUser(c echo.Context) error {
	username := strings.TrimSpace(c.FormValue("username"))
	password := strings.TrimSpace(c.FormValue("password"))

	if len(username) < 1 || len(username) > 32 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error_code": "create_failed",
			"error":      "Failed to create user",
		})
	}

	err := createUser(username, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error_code": "create_failed",
			"error":      "Failed to create user",
		})
	}

	return c.JSON(http.StatusOK, "")
}

// Saves new user into the database, first hashing their password
func createUser(username string, password string) error {
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`
			INSERT INTO site_user (username, password) VALUES ($1, $2);
		`,
		username, hash,
	)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// Wrapper around bcrypt's function to hash password
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
