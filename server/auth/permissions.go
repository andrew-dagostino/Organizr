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
				SELECT board_member.id
				FROM board_member
				JOIN board ON (board.id = board_member.board_id)
				WHERE board_member.member_id = $1
				AND board.gid = $2
				AND board_member.board_permission_id <= $3
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
				SELECT board_member.id
				FROM board_member
				JOIN task_column ON (task_column.board_id = board_member.board_id)
				WHERE board_member.member_id = $1
				AND task_column.gid = $2
				AND board_member.board_permission_id <= $3
			) AS has_permission;
		`,
		memberId, columnGid, minPermission,
	).Scan(&hasPermission)

	if err != nil {
		return false, err
	}
	return hasPermission, nil
}

func VerifyTaskPermission(memberId int, taskGid string, minPermission int) (bool, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		return false, err
	}
	defer conn.Close(context.Background())

	hasPermission := false
	err = conn.QueryRow(context.Background(),
		`
			SELECT EXISTS (
				SELECT board_member.id
				FROM board_member
				JOIN task_column ON (task_column.board_id = board_member.board_id)
				JOIN task ON (task.task_column_id = task_column.id)
				WHERE board_member.member_id = $1
				AND task.gid = $2
				AND board_member.board_permission_id <= $3
			) AS has_permission;
		`,
		memberId, taskGid, minPermission,
	).Scan(&hasPermission)

	if err != nil {
		return false, err
	}
	return hasPermission, nil
}
