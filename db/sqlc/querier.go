// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
)

type Querier interface {
	CommentTrending(ctx context.Context, dateSUB interface{}) ([]*CommentTrendingRow, error)
	CommentTrendingUsers(ctx context.Context, dateSUB interface{}) ([]*CommentTrendingUsersRow, error)
	CreatePolls(ctx context.Context, arg CreatePollsParams) error
	CreateStories(ctx context.Context, arg CreateStoriesParams) error
	DeleteExpiredTokens(ctx context.Context) error
	DeleteTrendingPosts(ctx context.Context, dateSUB interface{}) error
	DeleteTrendingUsers(ctx context.Context, dateSUB interface{}) error
	GetAdmin(ctx context.Context, mailID string) ([]*AdminUser, error)
	GetFeedback(ctx context.Context) ([]*Feedback, error)
	GetFollowersCount(ctx context.Context) ([]int32, error)
	GetFollowing(ctx context.Context, dateSUB interface{}) ([]*GetFollowingRow, error)
	GetFollowingIds(ctx context.Context, followerID int32) ([]int32, error)
	GetFollowingReaction(ctx context.Context, followerID int32) ([]*GetFollowingReactionRow, error)
	GetFollwers(ctx context.Context, dateSUB interface{}) ([]*GetFollwersRow, error)
	GetMutual(ctx context.Context, followerID int32) ([]int32, error)
	GetPolls(ctx context.Context) ([]*Poll, error)
	GetPosts(ctx context.Context, dateSUB interface{}) ([]*GetPostsRow, error)
	GetRating(ctx context.Context, userID int32) (int32, error)
	GetStories(ctx context.Context) ([]*Story, error)
	GetUserComments(ctx context.Context, dateSUB interface{}) ([]*GetUserCommentsRow, error)
	GetUserLikes(ctx context.Context, dateSUB interface{}) ([]*GetUserLikesRow, error)
	GetUsers(ctx context.Context) (int64, error)
	InsertAdmin(ctx context.Context, arg InsertAdminParams) error
	InsertTrending(ctx context.Context, postID int64) error
	InsertTrendingUsers(ctx context.Context, userID int32) error
	LikeTrending(ctx context.Context, dateSUB interface{}) ([]*LikeTrendingRow, error)
	LikeTrendingUsers(ctx context.Context, dateSUB interface{}) ([]*LikeTrendingUsersRow, error)
	UpdateComments(ctx context.Context, arg UpdateCommentsParams) error
	UpdateRating(ctx context.Context, arg UpdateRatingParams) error
	UpsertPostRecommendations(ctx context.Context, arg UpsertPostRecommendationsParams) error
	UpsertUserRecommendations(ctx context.Context, arg UpsertUserRecommendationsParams) error
}

var _ Querier = (*Queries)(nil)
