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

func GetTasks(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	columnGid := c.Param("column_gid")

	hasPermission, err := auth.VerifyColumnPermission(memberId, columnGid, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "get_tasks_failed",
			"message": "Failed to retrieve tasks",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":    "invalid_permission",
			"message": "Invalid permissions to retrieve tasks",
		})
	}

	tasks, err := retrieveAllTasks(columnGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "get_tasks_failed",
			"message": "Failed to retrieve tasks",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	columnGid := c.Param("column_gid")
	taskGid := c.Param("task_gid")

	hasPermission, err := auth.VerifyColumnPermission(memberId, columnGid, auth.VIEW_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "get_task_failed",
			"message": "Failed to retrieve task",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":    "invalid_permission",
			"message": "Invalid permissions to retrieve task",
		})
	}

	task, err := retrieveTaskByGid(columnGid, taskGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "get_task_failed",
			"message": "Failed to retrieve task",
		})
	}

	return c.JSON(http.StatusOK, task)
}

func EditTask(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	title := strings.TrimSpace(c.FormValue("title"))
	description := strings.TrimSpace(c.FormValue("description"))

	columnGid := c.Param("column_gid")
	taskGid := c.Param("task_gid")

	hasPermission, err := auth.VerifyColumnPermission(memberId, columnGid, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "update_task_failed",
			"message": "Failed to update task",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":    "invalid_permission",
			"message": "Invalid permissions to update task",
		})
	}

	task, err := updateTask(columnGid, taskGid, title, description)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "update_task_failed",
			"message": "Failed to update task",
		})
	}

	return c.JSON(http.StatusOK, task)
}

func CreateTask(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	title := strings.TrimSpace(c.FormValue("title"))
	description := strings.TrimSpace(c.FormValue("description"))

	columnGid := c.Param("column_gid")

	hasPermission, err := auth.VerifyColumnPermission(memberId, columnGid, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "update_task_failed",
			"message": "Failed to update task",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":    "invalid_permission",
			"message": "Invalid permissions to create task",
		})
	}

	task, err := addTask(columnGid, title, description)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "add_task_failed",
			"message": "Failed to create new task",
		})
	}

	return c.JSON(http.StatusCreated, task)
}

func DeleteTask(c echo.Context, log *log.Logger) error {
	member := c.Get("user").(*jwt.Token)
	claims := member.Claims.(jwt.MapClaims)
	memberId := int(claims["id"].(float64))

	taskGid := c.Param("task_gid")
	columnGid := c.Param("column_gid")

	hasPermission, err := auth.VerifyColumnPermission(memberId, columnGid, auth.EDIT_PERM)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "delete_task_failed",
			"message": "Failed to delete task",
		})
	}

	if !hasPermission {
		return c.JSON(http.StatusForbidden, map[string]string{
			"code":    "invalid_permission",
			"message": "Invalid permissions to delete task",
		})
	}

	err = removeTask(taskGid)
	if err != nil {
		log.Error(strings.TrimSpace(err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "delete_task_failed",
			"message": "Failed to delete task",
		})
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

func updateTask(columnGid string, taskGid string, title string, description string) (types.Task, error) {
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
		title, description, taskGid, columnGid,
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
		taskGid,
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

func addTask(columnGid string, title string, description string) (types.Task, error) {
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
		columnGid, title, description,
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
