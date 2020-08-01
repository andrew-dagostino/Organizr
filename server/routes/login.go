package routes

import (
	"context"
	"net/http"
	"os"
	"strings"
	"test-website/server/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Authenticates a user with their username and password from a POST request
// and returns a new JWT session token
func LoginUser(c echo.Context) error {
	username := strings.TrimSpace(c.FormValue("username"))
	password := strings.TrimSpace(c.FormValue("password"))

	user, err := getUser(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "login_failed",
			"error": "Failed to log in user",
		})
	}

	success, err := verifyUser(user.Username, password)
	if success == false || err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"code":  "verif_error",
			"error": "Username and/or password are incorrect",
		})
	}

	token, err := generateJWT(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":  "login_failed",
			"error": "Failed to log in user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

// Creates a User struct from the username's assosciated user data
func getUser(username string) (types.User, error) {
	var u types.User

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return u, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(),
		`SELECT id, username, last_login
		FROM site_user
		WHERE username = $1;`,
		username,
	).Scan(&u.Id, &u.Username, &u.Last_Login)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Verifies that the username and password are correct
func verifyUser(username string, password string) (bool, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return false, err
	}
	defer conn.Close(context.Background())

	var hashed_password string
	err = conn.QueryRow(context.Background(),
		`
			SELECT password
			FROM site_user
			WHERE username = $1;
		`,
		username,
	).Scan(&hashed_password)
	if err != nil {
		return false, err
	}

	return comparePasswords(hashed_password, password), nil
}

// Wrapper around a bcrypt function to compare passwords
func comparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}

// Generates a new JWT using the user's information
func generateJWT(user types.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
