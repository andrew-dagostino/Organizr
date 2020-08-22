package routes

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"
	"test-website/server/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
)

func GetPosts(c echo.Context) error {
	posts, err := retrievePosts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_failed",
			"error": "Failed to get posts",
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func GetPostById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_failed",
			"error": "Failed to get posts",
		})
	}

	post, err := retrievePostById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_failed",
			"error": "Failed to get user info",
		})
	}

	return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))

	title := strings.TrimSpace(c.FormValue("title"))
	text := strings.TrimSpace(c.FormValue("text"))

	postId, err := addPost(userId, title, text)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "post_create_failed",
			"error": "Failed to create new post",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"id": string(postId),
	})
}

func retrievePosts() ([]types.Post, error) {
	posts := []types.Post{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return posts, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT id, title, text, created, updated
		FROM post;`,
	)
	if err != nil {
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		var post types.Post
		err = rows.Scan(&post.Id, &post.Title, &post.Text, &post.Created, &post.Updated)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func retrievePostById(id int) (types.Post, error) {
	post := types.Post{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return post, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(),
		`
		SELECT id, title, text, created, updated
		FROM post
		WHERE id = $1;
		`,
		id,
	).Scan(&post.Id, &post.Title, &post.Text, &post.Created, &post.Updated)
	if err != nil {
		return post, err
	}

	return post, nil
}

func addPost(userId int, title string, text string) (int, error) {
	postId := -1

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return postId, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return postId, err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(),
		`
		INSERT INTO post (author, title, text)
		VALUES ($1, $2, $3)
		RETURNING id;
		`,
		userId, title, text,
	).Scan(&postId)
	if err != nil {
		return -1, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return -1, err
	}

	return postId, nil
}
