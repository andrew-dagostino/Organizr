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
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// swagger:parameters column column-retrieve-all
type GetColumnsRequest struct {
	// UUID of parent board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`
}

// swagger:response multi-column-response
type GetColumnsResponse struct {
	// in: body
	Body []types.TaskColumn
}

// swagger:parameters column column-retrieve-one
type GetColumnRequest struct {
	// UUID of parent board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`

	// UUID of column
	//
	// in: path
	// required: true
	Column_GID string `param:"column_gid"`
}

// swagger:response single-column-response
type GetColumnResponse struct {
	// in: body
	Body types.TaskColumn
}

// swagger:parameters column column-update
type UpdateColumnRequest struct {
	// UUID of parent board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`

	// UUID of column
	//
	// in: path
	// required: true
	Column_GID string `param:"column_gid"`

	// Title of column
	//
	// in: body
	Title string `form:"title"`
}

// swagger:parameters column column-delete
type DeleteColumnRequest struct {
	// UUID of parent board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`

	// UUID of column
	//
	// in: path
	// required: true
	Column_GID string `param:"column_gid"`
}

// swagger:route GET /api/column/{Board_GID} column column-retrieve-all
//
// Retrieves all columns by parent board UUID
//
// Security:
// - Bearer: []
//
// Responses:
//   200: multi-column-response
//   400: error-response
func GetColumns(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Body.Code = "get_columns_failed"
	e.Body.Message = "Failed to retrieve columns"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(GetColumnsRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	hasPermission, err := auth.VerifyBoardPermission(memberId, params.Board_GID, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Body.Code = "invalid_permission"
		e.Body.Message = "Invalid permissions to retrieve columns"
		return c.JSON(http.StatusForbidden, e)
	}

	columns, err := retrieveAllColumns(params.Board_GID)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, columns)
}

// swagger:route GET /api/column/{Board_GID}/{Column_GID} column column-retrieve-one
//
// Retrieves column by parent board and column UUIDs
//
// Security:
// - Bearer: []
//
// Responses:
//   200: single-column-response
//   400: error-response
func GetColumnById(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Body.Code = "get_column_failed"
	e.Body.Message = "Failed to retrieve column"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(GetColumnRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	hasPermission, err := auth.VerifyBoardPermission(memberId, params.Board_GID, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Body.Code = "invalid_permission"
		e.Body.Message = "Invalid permissions to retrieve column"
		return c.JSON(http.StatusForbidden, e)
	}

	column, err := retrieveColumnByGid(params.Board_GID, params.Column_GID)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, column)
}

// swagger:route PUT /api/column/{Board_GID}/{Column_GID} column column-update
//
// Updates column by parent board and column UUIDs
//
// Security:
// - Bearer: []
//
// Responses:
//   200: single-column-response
//   400: error-response
func EditColumn(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Body.Code = "update_task_failed"
	e.Body.Message = "Failed to update column"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(UpdateColumnRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}
	cleanColumnData(params)

	hasPermission, err := auth.VerifyBoardPermission(memberId, params.Board_GID, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Body.Code = "invalid_permission"
		e.Body.Message = "Invalid permissions to update column"
		return c.JSON(http.StatusForbidden, e)
	}

	column, err := updateColumn(params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, column)
}

// swagger:route POST /api/column/{Board_GID} column column-create
//
// Creates a column in the board specified by UUID
//
// Security:
// - Bearer: []
//
// Responses:
//   200: single-column-response
//   400: error-response
func CreateColumn(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Body.Code = "add_column_failed"
	e.Body.Message = "Failed to create new column"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(UpdateColumnRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}
	cleanColumnData(params)

	hasPermission, err := auth.VerifyBoardPermission(memberId, params.Board_GID, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Body.Code = "invalid_permission"
		e.Body.Message = "Invalid permissions to create column"
		return c.JSON(http.StatusForbidden, e)
	}

	column, err := addColumn(params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusCreated, column)
}

// swagger:route DELETE /api/column/{Board_GID}/{Column_GID} column column-delete
//
// Deletes column by parent board and column UUIDs
//
// Security:
// - Bearer: []
//
// Responses:
//   200:
//   400: error-response
func DeleteColumn(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Body.Code = "delete_column_failed"
	e.Body.Message = "Failed to delete column"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(DeleteColumnRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	hasPermission, err := auth.VerifyBoardPermission(memberId, params.Board_GID, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Body.Code = "invalid_permission"
		e.Body.Message = "Invalid permissions to delete column"
		return c.JSON(http.StatusForbidden, e)
	}

	err = removeColumn(params.Column_GID)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
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

func updateColumn(params *UpdateColumnRequest) (types.TaskColumn, error) {
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
		params.Title, params.Column_GID,
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
		params.Column_GID,
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

func addColumn(params *UpdateColumnRequest) (types.TaskColumn, error) {
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
		params.Board_GID, params.Title,
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

// Removes whitespace from title
func cleanColumnData(params *UpdateColumnRequest) {
	params.Title = strings.TrimSpace(params.Title)
}
