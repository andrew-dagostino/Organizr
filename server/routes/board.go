package routes

import (
	"context"
	"net/http"
	"os"
	"strings"
	"test-website/server/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
)

func GetBoards(c echo.Context) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberGid := claims["gid"].(string)

	boards, err := retrieveAllBoards(memberGid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_failed",
			"error": "Failed to get posts",
		})
	}

	return c.JSON(http.StatusOK, boards)
}

func GetBoardById(c echo.Context) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberGid := claims["gid"].(string)

	boardGid := c.Param("board_gid")

	board, err := retrieveBoardByGid(memberGid, boardGid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_failed",
			"error": "Failed to get user info",
		})
	}

	return c.JSON(http.StatusOK, board)
}

func CreateBoard(c echo.Context) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberGid := claims["gid"].(string)

	title := strings.TrimSpace(c.FormValue("title"))

	boardGid, err := addBoard(memberGid, title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "post_create_failed",
			"error": "Failed to create new post",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"id": string(boardGid),
	})
}

func retrieveAllBoards(memberGid string) ([]types.Board, error) {
	boards := []types.Board{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return boards, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`
			SELECT
				id,
				gid,
				title
			FROM board
			JOIN board_member ON (board_member.board_id = board.id)
			JOIN member ON (member.id = board_member.member_id)
			WHERE member.gid = $1;
		`,
		memberGid,
	)
	if err != nil {
		return boards, err
	}

	defer rows.Close()

	for rows.Next() {
		var board types.Board
		err = rows.Scan(&board.Id, &board.Gid, &board.Title)
		if err != nil {
			return boards, err
		}
		boards = append(boards, board)
	}

	return boards, nil
}

func retrieveBoardByGid(memberGid string, boardGid string) (types.Board, error) {
	board := types.Board{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return board, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title
			FROM board
			JOIN board_member ON (board_member.board_id = board.id)
			JOIN member ON (member.id = board_member.member_id)
			WHERE member.gid = $1 AND board.gid = $2;
		`,
		memberGid, boardGid,
	).Scan(&board.Id, &board.Gid, &board.Title)

	if err != nil {
		return board, err
	}
	return board, nil
}

func addBoard(memberGid string, title string) (string, error) {
	boardGid := ""

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return boardGid, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return boardGid, err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(),
		`
			INSERT INTO board(title)
			VALUES ($1)
			RETURNING gid;
		`,
		title,
	).Scan(&boardGid)
	if err != nil {
		return "", err
	}

	owner_permission := 1
	_, err = tx.Exec(context.Background(),
		`
			WITH new_board AS (
				SELECT id
				FROM board
				WHERE gid = $2
			)

			INSERT INTO board_member(member_id, board_id, permission_id)
			VALUES ($1, SELECT id FROM new_board, $3);
		`,
		title, boardGid, owner_permission,
	)
	if err != nil {
		return "", err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return "", err
	}

	return boardGid, nil
}
