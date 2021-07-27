package models

// swagger:parameters authentication login
type LoginRequest struct {
	// in: formData
	// required: true
	// example: user@email.com
	Username string `form:"username" json:"username"`

	// in: formData
	// required: true
	// example: password1234
	Password string `form:"password" json:"password"`
}

// swagger:parameters authentication register
type RegisterRequest struct {
	// in: formData
	// required: true
	// example: andrew
	Username string `form:"username" json:"username"`

	// in: formData
	// required: true
	// example: user@email.com
	Email string `form:"email" json:"email"`

	// in: formData
	// required: true
	// example: password1234
	Password string `form:"password" json:"password"`
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
	Title string `form:"title" json:"title"`
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
	// in: query
	// required: true
	Board_GID string `query:"board_gid"`
}

// swagger:parameters column column-retrieve-one
type GetColumnRequest struct {
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
	// in: body
	// required: true
	Board_GID string `form:"board_gid" json:"board_gid"`

	// UUID of column
	//
	// in: path
	// required: true
	Column_GID string `param:"column_gid"`

	// Title of column
	//
	// in: body
	Title string `form:"title" json:"title"`
}

// swagger:parameters column column-delete
type DeleteColumnRequest struct {
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
	// in: query
	// required: true
	Column_GID string `query:"column_gid"`
}

// swagger:parameters task task-retrieve-one
type GetTaskRequest struct {
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
	// in: body
	// required: true
	Column_GID string `form:"column_gid" json:"column_gid"`

	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string `param:"task_gid"`

	// Title of task
	//
	// in: body
	Title string `form:"title" json:"title"`

	// Description of task
	//
	// in: body
	Description string `form:"description" json:"description"`
}

// swagger:parameters task task-delete
type DeleteTaskRequest struct {
	// UUID of task
	//
	// in: path
	// required: true
	Task_GID string `param:"task_gid"`
}
