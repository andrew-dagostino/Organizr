package types

type Member struct {
	Id       int    `json:"id"`
	Gid      string `json:"gid"`
	Username string `json:"username"`
}

type Board struct {
	Id          int    `json:"id"`
	Gid         string `json:"gid"`
	Title       string `json:"title"`
	MemberCount int    `json:"board_member_count"`
}

type TaskColumn struct {
	Id      int    `json:"id"`
	Gid     string `json:"gid"`
	Title   string `json:"title"`
	BoardId int    `json:"board_id"`
}

type Task struct {
	Id           int    `json:"id"`
	Gid          string `json:"gid"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	TaskColumnId int    `json:"task_column_id"`
}
