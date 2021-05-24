package routes

import (
	"context"
	"net/http"
	"organizr/server/auth"
	"organizr/server/types"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func GetBoards(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberGid := claims["gid"].(string)

	boards, err := retrieveAllBoards(memberGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_boards_failed",
			"error": "Failed to retrieve boards",
		})
	}

	return c.JSON(http.StatusOK, boards)
}

func GetBoardById(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))
	memberGid := claims["gid"].(string)

	boardGid := c.Param("board_gid")

	hasPermission, err := auth.VerifyBoardPermission(memberId, boardGid, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "update_board_failed",
			"error": "Failed to update board",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to update board",
		})
	}

	board, err := retrieveBoardByGid(memberGid, boardGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_board_failed",
			"error": "Failed to retrieve board",
		})
	}

	return c.JSON(http.StatusOK, board)
}

func EditBoard(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	title := strings.TrimSpace(c.FormValue("title"))
	boardGid := c.Param("board_gid")

	hasPermission, err := auth.VerifyBoardPermission(memberId, boardGid, auth.OWNER_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "update_board_failed",
			"error": "Failed to update board",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to update board",
		})
	}

	board, err := updateBoard(boardGid, title)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "update_board_failed",
			"error": "Failed to update board",
		})
	}

	return c.JSON(http.StatusOK, board)
}

func CreateBoard(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	title := strings.TrimSpace(c.FormValue("title"))

	board, err := addBoard(memberId, title)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "add_board_failed",
			"error": "Failed to create new board",
		})
	}

	return c.JSON(http.StatusCreated, board)
}

func DeleteBoard(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	boardGid := c.Param("board_gid")

	hasPermission, err := auth.VerifyBoardPermission(memberId, boardGid, auth.OWNER_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "delete_board_failed",
			"error": "Failed to delete board",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to delete board",
		})
	}

	err = removeBoard(boardGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "delete_board_failed",
			"error": "Failed to delete board",
		})
	}

	return c.JSON(http.StatusAccepted, nil)
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
				board.id,
				board.gid,
				board.title
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
				board.id,
				board.gid,
				board.title
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

func updateBoard(boardGid string, title string) (types.Board, error) {
	var board types.Board

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return board, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return board, err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`
			UPDATE board
			SET
				title = $1,
				updated = CURRENT_TIMESTAMP
			WHERE gid = $2;
		`,
		title, boardGid,
	)
	if err != nil {
		return board, err
	}

	err = tx.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title
			FROM board
			WHERE gid = $1;
		`,
		boardGid,
	).Scan(&board.Id, &board.Gid, &board.Title)
	if err != nil {
		return board, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return board, err
	}

	return board, nil
}

func addBoard(memberId int, title string) (types.Board, error) {
	var board types.Board

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return board, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return board, err
	}
	defer tx.Rollback(context.Background())

	boardId := -1
	err = tx.QueryRow(context.Background(),
		`
			INSERT INTO board(title)
			VALUES ($1)
			RETURNING id;
		`,
		title,
	).Scan(&boardId)
	if err != nil {
		return board, err
	}

	ownerPermission := 1
	_, err = tx.Exec(context.Background(),
		`
			INSERT INTO board_member(member_id, board_id, board_permission_id)
			VALUES ($1, $2, $3);
		`,
		memberId, boardId, ownerPermission,
	)
	if err != nil {
		return board, err
	}

	err = tx.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title
			FROM board
			WHERE id = $1;
		`,
		boardId,
	).Scan(&board.Id, &board.Gid, &board.Title)
	if err != nil {
		return board, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return board, err
	}

	return board, nil
}

func removeBoard(boardGid string) error {
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
			DELETE FROM board
			WHERE gid = $1;
		`,
		boardGid,
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
