package types

import "time"

type User struct {
	Id         int        `json:"id"`
	Username   string     `json:"username"`
	Last_Login *time.Time `json:"last_login"`
}
