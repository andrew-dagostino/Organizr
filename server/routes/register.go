package routes

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// Registers a new member using the POST's username and password
func RegisterMember(c echo.Context, log *log.Logger) error {
	username := strings.ToLower(strings.TrimSpace(c.FormValue("username")))
	email := strings.ToLower(strings.TrimSpace(c.FormValue("email")))
	password := strings.TrimSpace(c.FormValue("password"))

	err := validateRegisterData(username, email, password)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error_code": "register_failed",
			"error":      "Failed to create member",
		})
	}

	err = createMember(username, email, password)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error_code": "register_failed",
			"error":      "Failed to create member",
		})
	}

	return c.JSON(http.StatusOK, "")
}

// Saves new member into the database, first hashing their password
func createMember(username string, email string, password string) error {
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
			INSERT INTO member (username, email, password)
			VALUES ($1, $2, $3);
		`,
		username, email, hash,
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

// Validates user data before insert
func validateRegisterData(username string, email string, password string) error {

	if len(username) < 1 || len(username) > 32 {
		return errors.New("Invalid Username")
	}

	if len(password) < 8 {
		return errors.New("Invalid Password")
	}

	return nil
}
