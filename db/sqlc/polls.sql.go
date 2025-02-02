// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: polls.sql

package db

import (
	"context"
)

const createPolls = `-- name: CreatePolls :exec
INSERT INTO polls (poll_by,poll_question,options_count,options)
VALUES (?,?,?,?)
`

type CreatePollsParams struct {
	PollBy       string `json:"poll_by"`
	PollQuestion string `json:"poll_question"`
	OptionsCount int32  `json:"options_count"`
	Options      string `json:"options"`
}

func (q *Queries) CreatePolls(ctx context.Context, arg CreatePollsParams) error {
	_, err := q.exec(ctx, q.createPollsStmt, createPolls,
		arg.PollBy,
		arg.PollQuestion,
		arg.OptionsCount,
		arg.Options,
	)
	return err
}

const getPolls = `-- name: GetPolls :many
SELECT poll_id, poll_by, poll_question, options_count, options, created_at FROM polls 
ORDER BY created_at DESC
`

func (q *Queries) GetPolls(ctx context.Context) ([]*Poll, error) {
	rows, err := q.query(ctx, q.getPollsStmt, getPolls)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Poll{}
	for rows.Next() {
		var i Poll
		if err := rows.Scan(
			&i.PollID,
			&i.PollBy,
			&i.PollQuestion,
			&i.OptionsCount,
			&i.Options,
			&i.CreatedAt,
		); err != nil {
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
