package auth

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

const (
	OWNER_PERM = iota + 1
	EDIT_PERM
	VIEW_PERM
)

func VerifyBoardPermission(memberId int, boardGid string, minPermission int) (bool, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return false, err
	}
	defer conn.Close(context.Background())

	hasPermission := false
	err = conn.QueryRow(context.Background(),
		`
			SELECT EXISTS(
				SELECT id
				FROM board_member
				WHERE member_id = $1
				AND board_id = (
					SELECT id FROM board WHERE gid = $2
				)
				AND board_permission_id <= $3
			) AS has_permission;
		`,
		memberId, boardGid, minPermission,
	).Scan(&hasPermission)

	if err != nil {
		return false, err
	}
	return hasPermission, nil
}

func VerifyColumnPermission(memberId int, columnGid string, minPermission int) (bool, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return false, err
	}
	defer conn.Close(context.Background())

	hasPermission := false
	err = conn.QueryRow(context.Background(),
		`
			SELECT EXISTS (
				SELECT id
				FROM board_member
				WHERE member_id = $1
				AND board_id = (
					SELECT board_id FROM task_column WHERE gid = $2
				)
				AND board_permission_id <= $3
			) AS has_permission;
		`,
		memberId, columnGid, minPermission,
	).Scan(&hasPermission)

	if err != nil {
		return false, err
	}
	return hasPermission, nil
}
