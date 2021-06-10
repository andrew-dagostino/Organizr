package routes

import (
	"context"
	"errors"
	"net/http"
	"organizr/server/types"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// swagger:parameters authentication register
type RegisterBodyParams struct {
	// in: formData
	// required: true
	// example: andrew
	Username string `form:"username"`

	// in: formData
	// required: true
	// example: user@email.com
	Email string `form:"email"`

	// in: formData
	// required: true
	// example: password1234
	Password string `form:"password"`
}

// swagger:route POST /api/register authentication register
//
// Registers a new member using the supplied username, email, and password
//
// Responses:
//   200:
//   400: error-response
func RegisterMember(c echo.Context, log *log.Logger) error {
	e := new(types.Error)
	e.Code = "register_failed"
	e.Message = "Failed to create member"

	params := new(RegisterBodyParams)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusInternalServerError, e)
	}
	cleanRegisterData(params)

	err := validateRegisterData(params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	err = createMember(params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, "")
}

// Saves new member into the database, first hashing their password
func createMember(params *RegisterBodyParams) error {
	hash, err := hashPassword(params.Password)
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
		params.Username, params.Email, hash,
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
func validateRegisterData(params *RegisterBodyParams) error {

	if len(params.Username) < 1 || len(params.Username) > 32 {
		return errors.New("Invalid Username")
	}

	if len(params.Password) < 8 {
		return errors.New("Invalid Password")
	}

	return nil
}

// Removes whitespace and lowers username, email
func cleanRegisterData(params *RegisterBodyParams) {
	params.Username = strings.ToLower(strings.TrimSpace(params.Username))
	params.Email = strings.ToLower(strings.TrimSpace(params.Email))
	params.Password = strings.TrimSpace(params.Password)
}
