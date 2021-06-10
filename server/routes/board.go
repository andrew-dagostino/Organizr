package routes

import (
	"context"
	"net/http"
	"organizr/server/auth"
	"organizr/server/models"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// swagger:route GET /api/board board board-retrieve-all
//
// Retrieves all boards
//
// Security:
// - Bearer: []
//
// Responses:
//   200: multi-board-response
//   400: error-response
func GetBoards(c echo.Context, log *log.Logger) error {
	e := new(models.Error)
	e.Code = "get_boards_failed"
	e.Message = "Failed to retrieve boards"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberGid := claims["gid"].(string)

	boards, err := retrieveAllBoards(memberGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, boards)
}

// swagger:route GET /api/board/{Board_GID} board board-retrieve-one
//
// Retrieves board by UUID
//
// Security:
// - Bearer: []
//
// Responses:
//   200: single-board-response
//   400: error-response
func GetBoardById(c echo.Context, log *log.Logger) error {
	e := new(models.Error)
	e.Code = "get_board_failed"
	e.Message = "Failed to retrieve board"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(models.GetBoardRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	hasPermission, err := auth.VerifyBoardPermission(memberId, params.Board_GID, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Code = "invalid_permission"
		e.Message = "Invalid permissions to retrieve board"
		return c.JSON(http.StatusForbidden, e)
	}

	board, err := retrieveBoardByGid(params.Board_GID)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, board)
}

// swagger:route PUT /api/board/{Board_GID} board board-update
//
// Updates board by UUID
//
// Security:
// - Bearer: []
//
// Responses:
//   200: single-board-response
//   400: error-response
func EditBoard(c echo.Context, log *log.Logger) error {
	e := new(models.Error)
	e.Code = "update_board_failed"
	e.Message = "Failed to update board"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(models.UpdateBoardRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}
	cleanBoardData(params)

	hasPermission, err := auth.VerifyBoardPermission(memberId, params.Board_GID, auth.OWNER_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Code = "invalid_permission"
		e.Message = "Invalid permissions to update board"
		return c.JSON(http.StatusForbidden, e)
	}

	board, err := updateBoard(params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, board)
}

// swagger:route POST /api/board board board-create
//
// Creates a board
//
// Security:
// - Bearer: []
//
// Responses:
//   200: single-board-response
//   400: error-response
func CreateBoard(c echo.Context, log *log.Logger) error {
	e := new(models.Error)
	e.Code = "add_board_failed"
	e.Message = "Failed to create new board"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(models.UpdateBoardRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}
	cleanBoardData(params)

	board, err := addBoard(memberId, params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusCreated, board)
}

// swagger:route DELETE /api/board/{Board_GID} board board-delete
//
// Deletes board by UUID
//
// Security:
// - Bearer: []
//
// Responses:
//   200:
//   400: error-response
func DeleteBoard(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	boardGid := c.Param("board_gid")

	hasPermission, err := auth.VerifyBoardPermission(memberId, boardGid, auth.OWNER_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "delete_board_failed",
			"message": "Failed to delete board",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":    "invalid_permission",
			"message": "Invalid permissions to delete board",
		})
	}

	err = removeBoard(boardGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "delete_board_failed",
			"message": "Failed to delete board",
		})
	}

	return c.JSON(http.StatusAccepted, nil)
}

func retrieveAllBoards(memberGid string) ([]models.Board, error) {
	boards := []models.Board{}

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
				board.title,
				(SELECT COUNT(id) FROM board_member WHERE board_id = board.id) AS board_member_count
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
		var board models.Board
		err = rows.Scan(&board.Id, &board.Gid, &board.Title, &board.MemberCount)
		if err != nil {
			return boards, err
		}
		boards = append(boards, board)
	}

	return boards, nil
}

func retrieveBoardByGid(boardGid string) (models.Board, error) {
	board := models.Board{}

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
				board.title,
				(SELECT COUNT(id) FROM board_member WHERE board_id = board.id) AS board_member_count
			FROM board
			JOIN board_member ON (board_member.board_id = board.id)
			WHERE board.gid = $1;
		`,
		boardGid,
	).Scan(&board.Id, &board.Gid, &board.Title, &board.MemberCount)

	if err != nil {
		return board, err
	}
	return board, nil
}

func updateBoard(params *models.UpdateBoardRequest) (models.Board, error) {
	var board models.Board

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
		params.Title, params.Board_GID,
	)
	if err != nil {
		return board, err
	}

	err = tx.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				(SELECT COUNT(id) FROM board_member) AS board_member_count
			FROM board
			WHERE gid = $1;
		`,
		params.Board_GID,
	).Scan(&board.Id, &board.Gid, &board.Title, &board.MemberCount)
	if err != nil {
		return board, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return board, err
	}

	return board, nil
}

func addBoard(memberId int, params *models.UpdateBoardRequest) (models.Board, error) {
	var board models.Board

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
		params.Title,
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
				title,
				(SELECT COUNT(id) FROM board_member) AS board_member_count
			FROM board
			WHERE id = $1;
		`,
		boardId,
	).Scan(&board.Id, &board.Gid, &board.Title, &board.MemberCount)
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

// Removes whitespace from title
func cleanBoardData(params *models.UpdateBoardRequest) {
	params.Title = strings.TrimSpace(params.Title)
}
