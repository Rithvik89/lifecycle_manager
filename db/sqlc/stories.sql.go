// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: stories.sql

package db

import (
	"context"
)

const createStories = `-- name: CreateStories :exec
INSERT INTO polls (poll_by,poll_question,options_count,options)
VALUES (?,?,?,?)
`

type CreateStoriesParams struct {
	PollBy       string `json:"poll_by"`
	PollQuestion string `json:"poll_question"`
	OptionsCount int32  `json:"options_count"`
	Options      string `json:"options"`
}

func (q *Queries) CreateStories(ctx context.Context, arg CreateStoriesParams) error {
	_, err := q.exec(ctx, q.createStoriesStmt, createStories,
		arg.PollBy,
		arg.PollQuestion,
		arg.OptionsCount,
		arg.Options,
	)
	return err
}

const getStories = `-- name: GetStories :many
SELECT story_id, mail_id, content, media, created_at FROM stories
`

func (q *Queries) GetStories(ctx context.Context) ([]*Story, error) {
	rows, err := q.query(ctx, q.getStoriesStmt, getStories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Story{}
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.StoryID,
			&i.MailID,
			&i.Content,
			&i.Media,
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
