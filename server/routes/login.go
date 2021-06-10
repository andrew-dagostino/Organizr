package routes

import (
	"context"
	"net/http"
	"organizr/server/models"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// swagger:route POST /api/login authentication login
//
// Authenticates a member with their username and password from a POST, returning a new JWT session token
//
// Responses:
//   200: login-response
//   400: error-response
func LoginMember(c echo.Context, log *log.Logger) error {
	e := new(models.Error)
	e.Code = "login_failed"
	e.Message = "Username and/or password are incorrect"

	params := new(models.LoginRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	member, err := getMember(strings.ToLower(params.Username))
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	success, err := verifyMember(member.Username, params.Password)
	if success == false || err != nil {
		log.Info(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	token, err := generateJWT(member)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	res := new(models.AuthDetail)
	res.JWT = token
	return c.JSON(http.StatusOK, res)
}

// Creates a Member struct from the username's assosciated user data
func getMember(username string) (models.Member, error) {
	var member models.Member

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
func generateJWT(member models.Member) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = member.Username
	claims["gid"] = member.Gid
	claims["id"] = member.Id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
