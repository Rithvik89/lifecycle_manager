// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: auth.sql

package db

import (
	"context"
)

const getAdmin = `-- name: GetAdmin :many
select mail_id, open_to from admin_users
WHERE mail_id = (?)
`

func (q *Queries) GetAdmin(ctx context.Context, mailID string) ([]*AdminUser, error) {
	rows, err := q.query(ctx, q.getAdminStmt, getAdmin, mailID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*AdminUser{}
	for rows.Next() {
		var i AdminUser
		if err := rows.Scan(&i.MailID, &i.OpenTo); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAdmin = `-- name: InsertAdmin :exec
INSERT INTO admin_users (mail_id,open_to) VALUES (?,?)
`

type InsertAdminParams struct {
	MailID string `json:"mail_id"`
	OpenTo string `json:"open_to"`
}

func (q *Queries) InsertAdmin(ctx context.Context, arg InsertAdminParams) error {
	_, err := q.exec(ctx, q.insertAdminStmt, insertAdmin, arg.MailID, arg.OpenTo)
	return err
}