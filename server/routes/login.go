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
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// Authenticates a member with their username and password from a POST, returning a new JWT session token
func LoginMember(c echo.Context, log *log.Logger) error {
	username := strings.ToLower(strings.TrimSpace(c.FormValue("username")))
	password := strings.TrimSpace(c.FormValue("password"))

	member, err := getMember(username)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "login_failed",
			"error": "Username and/or password are incorrect",
		})
	}

	success, err := verifyMember(member.Username, password)
	if success == false || err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"code":  "login_failed",
			"error": "Username and/or password are incorrect",
		})
	}

	token, err := generateJWT(member)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":  "login_failed",
			"error": "Username and/or password are incorrect",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"jwt": token,
	})
}

// Creates a Member struct from the username's assosciated user data
func getMember(username string) (types.Member, error) {
	var member types.Member

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return member, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				username
			FROM member
			WHERE username = $1 OR email = $1;
		`,
		username,
	).Scan(&member.Id, &member.Gid, &member.Username)
	if err != nil {
		return member, err
	}

	return member, nil
}

// Verifies that the username and password are correct
func verifyMember(username string, password string) (bool, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return false, err
	}
	defer conn.Close(context.Background())

	var hashed_password string
	err = conn.QueryRow(context.Background(),
		`
			SELECT
				password
			FROM member
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
func generateJWT(member types.Member) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = member.Username
	claims["gid"] = member.Gid
	claims["id"] = member.Id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
