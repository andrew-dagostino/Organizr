package models

// swagger:parameters authentication login
type LoginRequest struct {
	// in: formData
	// required: true
	// example: user@email.com
	Username string `form:"username"`

	// in: formData
	// required: true
	// example: password1234
	Password string `form:"password"`
}

// swagger:parameters authentication register
type RegisterRequest struct {
	// in: formData
	// required: true
	// example: andrew
	Username string `form:"username"`

	// in: formData
	// required: true
	// example: user@email.com
	Email string `form:"email"`

	// in: formData
	// required: true
	// example: password1234
	Password string `form:"password"`
}

// swagger:parameters board board-retrieve-one
type GetBoardRequest struct {
	// UUID of board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`
}

// swagger:parameters board board-update
type UpdateBoardRequest struct {
	// UUID of board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`

	// Title of board
	//
	// in: body
	Title string `form:"title"`
}

// swagger:parameters board board-delete
type DeleteBoardRequest struct {
	// UUID of board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`
}

// swagger:parameters column column-retrieve-all
type GetColumnsRequest struct {
	// UUID of parent board
	//
	// in: path
	// required: true
	Board_GID string `param:"board_gid"`
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

// swagger:parameters task task-retrieve-all
type GetTasksRequest struct {
	// UUID of parent column
	//
	// in: path
	// required: true
	Column_GID string `param:"column_gid"`
}

// swagger:parameters task task-retrieve-one
type GetTaskRequest struct {
	// UUID of parent column
	//
	// in: path
	// required: true
	Column_GID string `param:"column_gid"`

	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string `param:"task_gid"`
}

// swagger:parameters task task-update
type UpdateTaskRequest struct {
	// UUID of parent column
	//
	// in: path
	// required: true
	Column_GID string `param:"column_gid"`

	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string `param:"task_gid"`

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
	Column_GID string `param:"column_gid"`

	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string `param:"task_gid"`
}
