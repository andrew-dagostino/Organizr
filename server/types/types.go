package types

import (
	"time"

	"github.com/jackc/pgtype"
)

type User struct {
	Id         int        `json:"id"`
	Username   string     `json:"username"`
	Last_Login *time.Time `json:"last_login"`
}

type Post struct {
	Id      int                `json:"id"`
	Title   string             `json:"title"`
	Text    string             `json:"text"`
	Img     *pgtype.ByteaArray `json:"img"`
	Created *time.Time         `json:"created"`
	Updated *time.Time         `json:"updated"`
}
