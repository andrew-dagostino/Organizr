package routes

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"test-website/server/types"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
)

// Authenticates a user with their username and password from a POST request
// and returns a new JWT session token
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_failed",
			"error": "Failed to get user info",
		})
	}

	user, err := retrieveUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_failed",
			"error": "Failed to get user info",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// Creates a User struct from the username's assosciated user data
func retrieveUser(id int) (types.User, error) {
	var user types.User

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return user, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(),
		`SELECT id, username, last_login
		FROM site_user
		WHERE id = $1;`,
		id,
	).Scan(&user.Id, &user.Username, &user.Last_Login)
	if err != nil {
		return user, err
	}

	return user, nil
}
