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

// swagger:parameters task task-retrieve-all
type GetTasksRequest struct {
	// UUID of parent column
	//
	// in: path
	// required: true
	Column_GID string
}

// swagger:response multi-task-response
type GetTasksResponse struct {
	// in: body
	Data []types.Task
}

// swagger:parameters task task-retrieve-one
type GetTaskRequest struct {
	// UUID of parent column
	//
	// in: path
	// required: true
	Column_GID string

	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string
}

// swagger:response single-task-response
type GetTaskResponse struct {
	// in: body
	Data types.Task
}

// swagger:parameters task task-update
type UpdateTaskRequest struct {
	// UUID of parent column
	//
	// in: path
	// required: true
	Column_GID string

	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string

	// Title of task
	//
	// in: body
	Title string `form:"title"`

	// Description of task
	//
	// in: body
	Description string `form:"description"`
}

// swagger:parameters task task-delete
type DeleteTaskRequest struct {
	// UUID of parent column
	//
	// in: path
	// required: true
	Column_GID string

	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string
}

// swagger:route GET /api/task/{Column_GID} task retrieve-all
//
// Retrieves all tasks by parent column UUID
//
// Responses:
//   200: multi-task-response
//   400: error-response
func GetTasks(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Data.Code = "get_tasks_failed"
	e.Data.Message = "Failed to retrieve tasks"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(GetTasksRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	hasPermission, err := auth.VerifyColumnPermission(memberId, params.Column_GID, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Data.Code = "invalid_permission"
		e.Data.Message = "Invalid permissions to retrieve tasks"
		return c.JSON(http.StatusForbidden, e)
	}

	tasks, err := retrieveAllTasks(params.Column_GID)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, tasks)
}

// swagger:route GET /api/task/{Column_GID}/{Task_GID} task retrieve-one
//
// Retrieves task by parent column and task UUIDs
//
// Responses:
//   200: single-task-response
//   400: error-response
func GetTaskById(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Data.Code = "get_task_failed"
	e.Data.Message = "Failed to retrieve task"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(GetTaskRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	hasPermission, err := auth.VerifyColumnPermission(memberId, params.Column_GID, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Data.Code = "invalid_permission"
		e.Data.Message = "Invalid permissions to retrieve task"
		return c.JSON(http.StatusForbidden, e)
	}

	task, err := retrieveTaskByGid(params.Column_GID, params.Task_GID)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, task)
}

// swagger:route PUT /api/task/{Column_GID}/{Task_GID} task update
//
// Updates task by parent column and task UUIDs
//
// Responses:
//   200: single-task-response
//   400: error-response
func EditTask(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Data.Code = "update_task_failed"
	e.Data.Message = "Failed to update task"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(UpdateTaskRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}
	cleanTaskData(params)

	hasPermission, err := auth.VerifyColumnPermission(memberId, params.Column_GID, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Data.Code = "invalid_permission"
		e.Data.Message = "Invalid permissions to update task"
		return c.JSON(http.StatusForbidden, e)
	}

	task, err := updateTask(params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, task)
}

// swagger:route POST /api/task/{Column_GID} task create
//
// Creates a task in the column specified by UUID
//
// Responses:
//   200: single-task-response
//   400: error-response
func CreateTask(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Data.Code = "add_task_failed"
	e.Data.Message = "Failed to create new task"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(UpdateTaskRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}
	cleanTaskData(params)

	hasPermission, err := auth.VerifyColumnPermission(memberId, params.Column_GID, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Data.Code = "invalid_permission"
		e.Data.Message = "Invalid permissions to create task"
		return c.JSON(http.StatusForbidden, e)
	}

	task, err := addTask(params)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusCreated, task)
}

// swagger:route DELETE /api/task/{Column_GID}/{Task_GID} task delete
//
// Deletes task by parent column and task UUIDs
//
// Responses:
//   200:
//   400: error-response
func DeleteTask(c echo.Context, log *log.Logger) error {
	e := new(Error)
	e.Data.Code = "delete_task_failed"
	e.Data.Message = "Failed to delete task"

	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	params := new(DeleteTaskRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	hasPermission, err := auth.VerifyColumnPermission(memberId, params.Column_GID, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	if !hasPermission {
		e.Data.Code = "invalid_permission"
		e.Data.Message = "Invalid permissions to delete task"
		return c.JSON(http.StatusForbidden, e)
	}

	err = removeTask(params.Task_GID)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusAccepted, nil)
}

func retrieveAllTasks(columnGid string) ([]types.Task, error) {
	tasks := []types.Task{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return tasks, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				description,
				task_column_id
			FROM task
			WHERE task_column_id = (
				SELECT id FROM task_column WHERE gid = $1
			);
		`,
		columnGid,
	)
	if err != nil {
		return tasks, err
	}

	defer rows.Close()

	for rows.Next() {
		var task types.Task
		err = rows.Scan(&task.Id, &task.Gid, &task.Title, &task.Description, &task.TaskColumnId)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func retrieveTaskByGid(columnGid string, taskGid string) (types.Task, error) {
	task := types.Task{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return task, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				description,
				task_column_id
			FROM task
			WHERE gid = $2
			AND task_column_id = (
				SELECT id FROM task_column WHERE gid = $1
			);
		`,
		columnGid, taskGid,
	).Scan(&task.Id, &task.Gid, &task.Title, &task.Description, &task.TaskColumnId)

	if err != nil {
		return task, err
	}
	return task, nil
}

func updateTask(params *UpdateTaskRequest) (types.Task, error) {
	var task types.Task

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return task, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return task, err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`
			UPDATE task
			SET
				title = $1,
				description = $2,
				task_column_id = (SELECT id FROM task_column WHERE gid = $4),
				updated = CURRENT_TIMESTAMP
			WHERE gid = $3;
		`,
		params.Title, params.Description, params.Task_GID, params.Column_GID,
	)
	if err != nil {
		return task, err
	}

	err = tx.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				description,
				task_column_id
			FROM task
			WHERE gid = $1;
		`,
		params.Task_GID,
	).Scan(&task.Id, &task.Gid, &task.Title, &task.Description, &task.TaskColumnId)
	if err != nil {
		return task, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return task, err
	}

	return task, nil
}

func addTask(params *UpdateTaskRequest) (types.Task, error) {
	var task types.Task

	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return task, err
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return task, err
	}
	defer tx.Rollback(context.Background())

	taskId := -1
	err = tx.QueryRow(context.Background(),
		`
			INSERT INTO task (title, description, task_column_id)
			SELECT $2, $3, id FROM task_column WHERE task_column.gid = $1
			RETURNING task.id;
		`,
		params.Column_GID, params.Title, params.Description,
	).Scan(&taskId)
	if err != nil {
		return task, err
	}

	err = tx.QueryRow(context.Background(),
		`
			SELECT
				id,
				gid,
				title,
				description,
				task_column_id
			FROM task
			WHERE id = $1;
		`,
		taskId,
	).Scan(&task.Id, &task.Gid, &task.Title, &task.Description, &task.TaskColumnId)
	if err != nil {
		return task, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return task, err
	}

	return task, nil
}

func removeTask(taskGid string) error {
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
			DELETE FROM task
			WHERE gid = $1;
		`,
		taskGid,
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

// Removes whitespace from title, description
func cleanTaskData(params *UpdateTaskRequest) {
	params.Title = strings.TrimSpace(params.Title)
	params.Description = strings.TrimSpace(params.Description)
}
