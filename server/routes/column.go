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

func GetColumns(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	boardGid := c.Param("board_gid")

	hasPermission, err := verifyBoardPermission(memberId, boardGid, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_columns_failed",
			"error": "Failed to retrieve columns",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to retrieve columns",
		})
	}

	columns, err := retrieveAllColumns(boardGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_columns_failed",
			"error": "Failed to retrieve columns",
		})
	}

	return c.JSON(http.StatusOK, columns)
}

func GetColumnById(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	boardGid := c.Param("board_gid")
	columnGid := c.Param("column_gid")

	hasPermission, err := verifyBoardPermission(memberId, boardGid, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_column_failed",
			"error": "Failed to retrieve column",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to retrieve column",
		})
	}

	column, err := retrieveColumnByGid(boardGid, columnGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "get_column_failed",
			"error": "Failed to retrieve column",
		})
	}

	return c.JSON(http.StatusOK, column)
}

func EditColumn(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	title := strings.TrimSpace(c.FormValue("title"))

	boardGid := c.Param("board_gid")
	columnGid := c.Param("column_gid")

	hasPermission, err := verifyBoardPermission(memberId, boardGid, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "update_column_failed",
			"error": "Failed to update column",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to update column",
		})
	}

	column, err := updateColumn(columnGid, title)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "update_column_failed",
			"error": "Failed to update column",
		})
	}

	return c.JSON(http.StatusOK, column)
}

func CreateColumn(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	title := strings.TrimSpace(c.FormValue("title"))

	boardGid := c.Param("board_gid")

	hasPermission, err := verifyBoardPermission(memberId, boardGid, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "update_column_failed",
			"error": "Failed to update column",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to create column",
		})
	}

	column, err := addColumn(boardGid, title)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "add_column_failed",
			"error": "Failed to create new column",
		})
	}

	return c.JSON(http.StatusCreated, column)
}

func DeleteColumn(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	columnGid := c.Param("column_gid")
	boardGid := c.Param("board_gid")

	hasPermission, err := verifyBoardPermission(memberId, boardGid, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "delete_column_failed",
			"error": "Failed to delete column",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":  "invalid_permission",
			"error": "Invalid permissions to delete column",
		})
	}

	err = removeColumn(columnGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":  "delete_column_failed",
			"error": "Failed to delete column",
		})
	}

	return c.JSON(http.StatusAccepted, nil)
}

func retrieveAllColumns(boardGid string) ([]types.TaskColumn, error) {
	columns := []types.TaskColumn{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return columns, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				board_id
			FROM task_column
			WHERE board_id = (
				SELECT id FROM board WHERE gid = $1
			);
		`,
		boardGid,
	)
	if err != nil {
		return columns, err
	}

	defer rows.Close()

	for rows.Next() {
		var column types.TaskColumn
		err = rows.Scan(&column.Id, &column.Gid, &column.Title, &column.BoardId)
		if err != nil {
			return columns, err
		}
		columns = append(columns, column)
	}

	return columns, nil
}

func retrieveColumnByGid(boardGid string, columnGid string) (types.TaskColumn, error) {
	column := types.TaskColumn{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return column, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				board_id
			FROM task_column
			WHERE gid = $2
			AND board_id = (
				SELECT id
				FROM board WHERE gid = $1
			);
		`,
		boardGid, columnGid,
	).Scan(&column.Id, &column.Gid, &column.Title, &column.BoardId)

	if err != nil {
		return column, err
	}
	return column, nil
}

func updateColumn(columnGid string, title string) (types.TaskColumn, error) {
	var column types.TaskColumn

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return column, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return column, err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`
			UPDATE task_column
			SET
				title = $1,
				updated = CURRENT_TIMESTAMP
			WHERE gid = $2;
		`,
		title, columnGid,
	)
	if err != nil {
		return column, err
	}

	err = tx.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				board_id
			FROM task_column
			WHERE gid = $1;
		`,
		columnGid,
	).Scan(&column.Id, &column.Gid, &column.Title, &column.BoardId)
	if err != nil {
		return column, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return column, err
	}

	return column, nil
}

func addColumn(boardGid string, title string) (types.TaskColumn, error) {
	var column types.TaskColumn

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return column, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return column, err
	}
	defer tx.Rollback(context.Background())

	columnId := -1
	err = tx.QueryRow(context.Background(),
		`
			INSERT INTO task_column (title, board_id)
			SELECT $2, id FROM board WHERE board.gid = $1
			RETURNING task_column.id;
		`,
		boardGid, title,
	).Scan(&columnId)
	if err != nil {
		return column, err
	}

	err = tx.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				board_id
			FROM task_column
			WHERE id = $1;
		`,
		columnId,
	).Scan(&column.Id, &column.Gid, &column.Title, &column.BoardId)
	if err != nil {
		return column, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return column, err
	}

	return column, nil
}

func removeColumn(columnGid string) error {
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
			DELETE FROM task_column
			WHERE gid = $1;
		`,
		columnGid,
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
