// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: rating.sql

package db

import (
	"context"
)

const getFollowersCount = `-- name: GetFollowersCount :many
SELECT followers from profiles
`

func (q *Queries) GetFollowersCount(ctx context.Context) ([]int32, error) {
	rows, err := q.query(ctx, q.getFollowersCountStmt, getFollowersCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var followers int32
		if err := rows.Scan(&followers); err != nil {
			return nil, err
		}
		items = append(items, followers)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowing = `-- name: GetFollowing :many
SELECT follower_id as user_id,count(following_id) from networks
WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
group BY follower_id
`

type GetFollowingRow struct {
	UserID int32 `json:"user_id"`
	Count  int64 `json:"count"`
}

func (q *Queries) GetFollowing(ctx context.Context) ([]*GetFollowingRow, error) {
	rows, err := q.query(ctx, q.getFollowingStmt, getFollowing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFollowingRow{}
	for rows.Next() {
		var i GetFollowingRow
		if err := rows.Scan(&i.UserID, &i.Count); err != nil {
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

const getFollwers = `-- name: GetFollwers :many
SELECT following_id as user_id,count(follower_id) from networks
WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
group by following_id
`

type GetFollwersRow struct {
	UserID int32 `json:"user_id"`
	Count  int64 `json:"count"`
}

func (q *Queries) GetFollwers(ctx context.Context) ([]*GetFollwersRow, error) {
	rows, err := q.query(ctx, q.getFollwersStmt, getFollwers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFollwersRow{}
	for rows.Next() {
		var i GetFollwersRow
		if err := rows.Scan(&i.UserID, &i.Count); err != nil {
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

const getPosts = `-- name: GetPosts :many
SELECT user_id,count(post_id) from posts
group by user_id
`

type GetPostsRow struct {
	UserID int32 `json:"user_id"`
	Count  int64 `json:"count"`
}

func (q *Queries) GetPosts(ctx context.Context) ([]*GetPostsRow, error) {
	rows, err := q.query(ctx, q.getPostsStmt, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetPostsRow{}
	for rows.Next() {
		var i GetPostsRow
		if err := rows.Scan(&i.UserID, &i.Count); err != nil {
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

const getRating = `-- name: GetRating :one
SELECT cric_index from profiles
WHERE user_id = (?)
`

func (q *Queries) GetRating(ctx context.Context, userID int32) (int32, error) {
	row := q.queryRow(ctx, q.getRatingStmt, getRating, userID)
	var cric_index int32
	err := row.Scan(&cric_index)
	return cric_index, err
}

const getUserComments = `-- name: GetUserComments :many
SELECT posts.user_id, count(parent_comments.post_id) as comment_count from parent_comments
join posts on posts.post_id = parent_comments.post_id
group by posts.user_id
`

type GetUserCommentsRow struct {
	UserID       int32 `json:"user_id"`
	CommentCount int64 `json:"comment_count"`
}

func (q *Queries) GetUserComments(ctx context.Context) ([]*GetUserCommentsRow, error) {
	rows, err := q.query(ctx, q.getUserCommentsStmt, getUserComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetUserCommentsRow{}
	for rows.Next() {
		var i GetUserCommentsRow
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

const getUserLikes = `-- name: GetUserLikes :many
SELECT posts.user_id, count(likes.post_id) as like_count from likes
join posts on posts.post_id = likes.post_id
group by posts.user_id
`

type GetUserLikesRow struct {
	UserID    int32 `json:"user_id"`
	LikeCount int64 `json:"like_count"`
}

func (q *Queries) GetUserLikes(ctx context.Context) ([]*GetUserLikesRow, error) {
	rows, err := q.query(ctx, q.getUserLikesStmt, getUserLikes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetUserLikesRow{}
	for rows.Next() {
		var i GetUserLikesRow
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

const updateRating = `-- name: UpdateRating :exec
UPDATE profiles 
SET cric_index = (?)
WHERE user_id = (?)
`

type UpdateRatingParams struct {
	CricIndex int32 `json:"cric_index"`
	UserID    int32 `json:"user_id"`
}

func (q *Queries) UpdateRating(ctx context.Context, arg UpdateRatingParams) error {
	_, err := q.exec(ctx, q.updateRatingStmt, updateRating, arg.CricIndex, arg.UserID)
	return err
}