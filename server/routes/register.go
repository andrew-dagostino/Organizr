package routes

import (
	"context"
	"net/http"
	"os"
	"strings"
	"test-website/server/types"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) (err error) {
	username := strings.TrimSpace(c.FormValue("username"))
	password := strings.TrimSpace(c.FormValue("password"))

	if len(username) < 1 || len(username) > 32 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error_code": "create_failed",
			"error":      "Failed to create user",
		})
	}

	user, err := createUser(username, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error_code": "create_failed",
			"error":      "Failed to create user",
		})
	}

	return c.JSON(http.StatusOK, user)
}

func createUser(username string, password string) (types.User, error) {
	var user types.User

	hash, err := hashPassword(password)
	if err != nil {
		return user, err
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return user, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return user, err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(),
		`INSERT INTO site_user (username, password) VALUES ($1, $2)
			RETURNING id, username;`,
		username, hash,
	).Scan(&user.Id, &user.Username)
	if err != nil {
		return user, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return user, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
