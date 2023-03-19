// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: trending_posts.sql

package db

import (
	"context"
)

const commentTrending = `-- name: CommentTrending :many
SELECT post_id,COUNT(user_id) as comment_count from parent_comments 
WHERE created_at >= DATE_SUB(NOW(),INTERVAL ? MINUTE) 
GROUP BY post_id
`

type CommentTrendingRow struct {
	PostID       int64 `json:"post_id"`
	CommentCount int64 `json:"comment_count"`
}

func (q *Queries) CommentTrending(ctx context.Context, dateSUB interface{}) ([]*CommentTrendingRow, error) {
	rows, err := q.query(ctx, q.commentTrendingStmt, commentTrending, dateSUB)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*CommentTrendingRow{}
	for rows.Next() {
		var i CommentTrendingRow
		if err := rows.Scan(&i.PostID, &i.CommentCount); err != nil {
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

const deleteTrendingPosts = `-- name: DeleteTrendingPosts :exec
DELETE from trending_posts 
where created_at < DATE_SUB(NOW(),INTERVAL ? MINUTE)
`

func (q *Queries) DeleteTrendingPosts(ctx context.Context, dateSUB interface{}) error {
	_, err := q.exec(ctx, q.deleteTrendingPostsStmt, deleteTrendingPosts, dateSUB)
	return err
}

const insertTrending = `-- name: InsertTrending :exec
INSERT INTO trending_posts (post_id) VALUES (?)
`

func (q *Queries) InsertTrending(ctx context.Context, postID int64) error {
	_, err := q.exec(ctx, q.insertTrendingStmt, insertTrending, postID)
	return err
}

const likeTrending = `-- name: LikeTrending :many
SELECT post_id,COUNT(user_id) as like_count from likes
WHERE created_at >= DATE_SUB(NOW(),INTERVAL ? MINUTE) 
GROUP BY post_id
`

type LikeTrendingRow struct {
	PostID    int64 `json:"post_id"`
	LikeCount int64 `json:"like_count"`
}

func (q *Queries) LikeTrending(ctx context.Context, dateSUB interface{}) ([]*LikeTrendingRow, error) {
	rows, err := q.query(ctx, q.likeTrendingStmt, likeTrending, dateSUB)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*LikeTrendingRow{}
	for rows.Next() {
		var i LikeTrendingRow
		if err := rows.Scan(&i.PostID, &i.LikeCount); err != nil {
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
