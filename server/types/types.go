package types

import "time"

type Error struct {
	Code  string `json:"error_code"`
	Error string `json:"error"`
}

type User struct {
	Id         int        `json:"id"`
	Username   string     `json:"username"`
	Last_Login *time.Time `json:"last_login"`
}
