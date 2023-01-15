// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: trending_users.sql

package db

import (
	"context"
)

const commentTrendingUsers = `-- name: CommentTrendingUsers :many
SELECT user_id,COUNT(post_id) as comment_count from parent_comments 
WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
GROUP BY user_id
`

type CommentTrendingUsersRow struct {
	UserID       int32 `json:"user_id"`
	CommentCount int64 `json:"comment_count"`
}

func (q *Queries) CommentTrendingUsers(ctx context.Context) ([]*CommentTrendingUsersRow, error) {
	rows, err := q.query(ctx, q.commentTrendingUsersStmt, commentTrendingUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*CommentTrendingUsersRow{}
	for rows.Next() {
		var i CommentTrendingUsersRow
		if err := rows.Scan(&i.UserID, &i.CommentCount); err != nil {
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

const deleteTrendingUsers = `-- name: DeleteTrendingUsers :exec
DELETE from trending_users
where create_at < DATE_SUB(NOW(),INTERVAL (?) DAY)
`

func (q *Queries) DeleteTrendingUsers(ctx context.Context, dateSUB interface{}) error {
	_, err := q.exec(ctx, q.deleteTrendingUsersStmt, deleteTrendingUsers, dateSUB)
	return err
}

const insertTrendingUsers = `-- name: InsertTrendingUsers :exec
INSERT INTO trending_users (user_id) VALUES (?)
`

func (q *Queries) InsertTrendingUsers(ctx context.Context, userID int32) error {
	_, err := q.exec(ctx, q.insertTrendingUsersStmt, insertTrendingUsers, userID)
	return err
}

const likeTrendingUsers = `-- name: LikeTrendingUsers :many
SELECT posts.user_id,COUNT(likes.post_id) as like_count from likes
INNER JOIN posts ON posts.post_id = likes.post_id
WHERE likes.created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
GROUP BY posts.user_id
`

type LikeTrendingUsersRow struct {
	UserID    int32 `json:"user_id"`
	LikeCount int64 `json:"like_count"`
}

func (q *Queries) LikeTrendingUsers(ctx context.Context) ([]*LikeTrendingUsersRow, error) {
	rows, err := q.query(ctx, q.likeTrendingUsersStmt, likeTrendingUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*LikeTrendingUsersRow{}
	for rows.Next() {
		var i LikeTrendingUsersRow
		if err := rows.Scan(&i.UserID, &i.LikeCount); err != nil {
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
