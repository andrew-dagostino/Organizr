package routes

import (
	"context"
	"net/http"
	"organizr/server/types"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// User's Authentication Information
//
// swagger:parameters authentication login
type LoginBodyParams struct {
	// in: formData
	// required: true
	// example: user@email.com
	Username string `form:"username"`

	// in: formData
	// required: true
	// example: password1234
	Password string `form:"password"`
}

// Authentication Response
//
// swagger:response login-response
type LoginBodyResponse struct {
	Data struct {
		// in: body
		// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
		JWT string `json:"jwt"`
	}
}

// Error Response
//
// swagger:response error-response
type Error struct {
	Data struct {
		// in: body
		// example: login_failed
		Code string `json:"code"`

		// in: body
		// example: Username and/or password are incorrect
		Message string `json:"message"`
	}
}

// swagger:route POST /api/login authentication login
//
// Authenticates a member with their username and password from a POST, returning a new JWT session token
//
// Responses:
//   200: login-response
//   400: error-response
func LoginMember(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Data.Code = "login_failed"
	e.Data.Message = "Username and/or password are incorrect"

	params := new(LoginBodyParams)
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

	res := new(LoginBodyResponse)
	res.Data.JWT = token
	return c.JSON(http.StatusOK, res)
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
