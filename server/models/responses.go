package models

// swagger:response error-response
type ErrorResponse struct {
	// in: body
	Body Error
}

// swagger:response login-response
type LoginBodyResponse struct {
	// in: body
	Body AuthDetail
}

// swagger:response multi-board-response
type GetBoardsResponse struct {
	// in: body
	Body []Board
}

// swagger:response single-board-response
type GetBoardResponse struct {
	// in: body
	Body Board
}

// swagger:response multi-column-response
type GetColumnsResponse struct {
	// in: body
	Body []TaskColumn
}

// swagger:response single-column-response
type GetColumnResponse struct {
	// in: body
	Body TaskColumn
}

// swagger:response multi-task-response
type GetTasksResponse struct {
	// in: body
	Body []Task
}

// swagger:response single-task-response
type GetTaskResponse struct {
	// in: body
	Body Task
}
