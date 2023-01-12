// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.commentTrendingStmt, err = db.PrepareContext(ctx, commentTrending); err != nil {
		return nil, fmt.Errorf("error preparing query CommentTrending: %w", err)
	}
	if q.commentTrendingUsersStmt, err = db.PrepareContext(ctx, commentTrendingUsers); err != nil {
		return nil, fmt.Errorf("error preparing query CommentTrendingUsers: %w", err)
	}
	if q.createPollsStmt, err = db.PrepareContext(ctx, createPolls); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePolls: %w", err)
	}
	if q.deleteExpiredTokensStmt, err = db.PrepareContext(ctx, deleteExpiredTokens); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExpiredTokens: %w", err)
	}
	if q.getAdminStmt, err = db.PrepareContext(ctx, getAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query GetAdmin: %w", err)
	}
	if q.getFeedbackStmt, err = db.PrepareContext(ctx, getFeedback); err != nil {
		return nil, fmt.Errorf("error preparing query GetFeedback: %w", err)
	}
	if q.getFollowersCountStmt, err = db.PrepareContext(ctx, getFollowersCount); err != nil {
		return nil, fmt.Errorf("error preparing query GetFollowersCount: %w", err)
	}
	if q.getFollowingStmt, err = db.PrepareContext(ctx, getFollowing); err != nil {
		return nil, fmt.Errorf("error preparing query GetFollowing: %w", err)
	}
	if q.getFollowingIdsStmt, err = db.PrepareContext(ctx, getFollowingIds); err != nil {
		return nil, fmt.Errorf("error preparing query GetFollowingIds: %w", err)
	}
	if q.getFollowingReactionStmt, err = db.PrepareContext(ctx, getFollowingReaction); err != nil {
		return nil, fmt.Errorf("error preparing query GetFollowingReaction: %w", err)
	}
	if q.getFollwersStmt, err = db.PrepareContext(ctx, getFollwers); err != nil {
		return nil, fmt.Errorf("error preparing query GetFollwers: %w", err)
	}
	if q.getMutualStmt, err = db.PrepareContext(ctx, getMutual); err != nil {
		return nil, fmt.Errorf("error preparing query GetMutual: %w", err)
	}
	if q.getPollsStmt, err = db.PrepareContext(ctx, getPolls); err != nil {
		return nil, fmt.Errorf("error preparing query GetPolls: %w", err)
	}
	if q.getPostsStmt, err = db.PrepareContext(ctx, getPosts); err != nil {
		return nil, fmt.Errorf("error preparing query GetPosts: %w", err)
	}
	if q.getRatingStmt, err = db.PrepareContext(ctx, getRating); err != nil {
		return nil, fmt.Errorf("error preparing query GetRating: %w", err)
	}
	if q.getUserCommentsStmt, err = db.PrepareContext(ctx, getUserComments); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserComments: %w", err)
	}
	if q.getUserLikesStmt, err = db.PrepareContext(ctx, getUserLikes); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserLikes: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.insertAdminStmt, err = db.PrepareContext(ctx, insertAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query InsertAdmin: %w", err)
	}
	if q.insertTrendingStmt, err = db.PrepareContext(ctx, insertTrending); err != nil {
		return nil, fmt.Errorf("error preparing query InsertTrending: %w", err)
	}
	if q.insertTrendingUsersStmt, err = db.PrepareContext(ctx, insertTrendingUsers); err != nil {
		return nil, fmt.Errorf("error preparing query InsertTrendingUsers: %w", err)
	}
	if q.likeTrendingStmt, err = db.PrepareContext(ctx, likeTrending); err != nil {
		return nil, fmt.Errorf("error preparing query LikeTrending: %w", err)
	}
	if q.likeTrendingUsersStmt, err = db.PrepareContext(ctx, likeTrendingUsers); err != nil {
		return nil, fmt.Errorf("error preparing query LikeTrendingUsers: %w", err)
	}
	if q.updateRatingStmt, err = db.PrepareContext(ctx, updateRating); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateRating: %w", err)
	}
	if q.upsertPostRecommendationsStmt, err = db.PrepareContext(ctx, upsertPostRecommendations); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertPostRecommendations: %w", err)
	}
	if q.upsertUserRecommendationsStmt, err = db.PrepareContext(ctx, upsertUserRecommendations); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertUserRecommendations: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.commentTrendingStmt != nil {
		if cerr := q.commentTrendingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing commentTrendingStmt: %w", cerr)
		}
	}
	if q.commentTrendingUsersStmt != nil {
		if cerr := q.commentTrendingUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing commentTrendingUsersStmt: %w", cerr)
		}
	}
	if q.createPollsStmt != nil {
		if cerr := q.createPollsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPollsStmt: %w", cerr)
		}
	}
	if q.deleteExpiredTokensStmt != nil {
		if cerr := q.deleteExpiredTokensStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExpiredTokensStmt: %w", cerr)
		}
	}
	if q.getAdminStmt != nil {
		if cerr := q.getAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAdminStmt: %w", cerr)
		}
	}
	if q.getFeedbackStmt != nil {
		if cerr := q.getFeedbackStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFeedbackStmt: %w", cerr)
		}
	}
	if q.getFollowersCountStmt != nil {
		if cerr := q.getFollowersCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFollowersCountStmt: %w", cerr)
		}
	}
	if q.getFollowingStmt != nil {
		if cerr := q.getFollowingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFollowingStmt: %w", cerr)
		}
	}
	if q.getFollowingIdsStmt != nil {
		if cerr := q.getFollowingIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFollowingIdsStmt: %w", cerr)
		}
	}
	if q.getFollowingReactionStmt != nil {
		if cerr := q.getFollowingReactionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFollowingReactionStmt: %w", cerr)
		}
	}
	if q.getFollwersStmt != nil {
		if cerr := q.getFollwersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFollwersStmt: %w", cerr)
		}
	}
	if q.getMutualStmt != nil {
		if cerr := q.getMutualStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMutualStmt: %w", cerr)
		}
	}
	if q.getPollsStmt != nil {
		if cerr := q.getPollsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPollsStmt: %w", cerr)
		}
	}
	if q.getPostsStmt != nil {
		if cerr := q.getPostsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsStmt: %w", cerr)
		}
	}
	if q.getRatingStmt != nil {
		if cerr := q.getRatingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRatingStmt: %w", cerr)
		}
	}
	if q.getUserCommentsStmt != nil {
		if cerr := q.getUserCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserCommentsStmt: %w", cerr)
		}
	}
	if q.getUserLikesStmt != nil {
		if cerr := q.getUserLikesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserLikesStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.insertAdminStmt != nil {
		if cerr := q.insertAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertAdminStmt: %w", cerr)
		}
	}
	if q.insertTrendingStmt != nil {
		if cerr := q.insertTrendingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertTrendingStmt: %w", cerr)
		}
	}
	if q.insertTrendingUsersStmt != nil {
		if cerr := q.insertTrendingUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertTrendingUsersStmt: %w", cerr)
		}
	}
	if q.likeTrendingStmt != nil {
		if cerr := q.likeTrendingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing likeTrendingStmt: %w", cerr)
		}
	}
	if q.likeTrendingUsersStmt != nil {
		if cerr := q.likeTrendingUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing likeTrendingUsersStmt: %w", cerr)
		}
	}
	if q.updateRatingStmt != nil {
		if cerr := q.updateRatingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateRatingStmt: %w", cerr)
		}
	}
	if q.upsertPostRecommendationsStmt != nil {
		if cerr := q.upsertPostRecommendationsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertPostRecommendationsStmt: %w", cerr)
		}
	}
	if q.upsertUserRecommendationsStmt != nil {
		if cerr := q.upsertUserRecommendationsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertUserRecommendationsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                            DBTX
	tx                            *sql.Tx
	commentTrendingStmt           *sql.Stmt
	commentTrendingUsersStmt      *sql.Stmt
	createPollsStmt               *sql.Stmt
	deleteExpiredTokensStmt       *sql.Stmt
	getAdminStmt                  *sql.Stmt
	getFeedbackStmt               *sql.Stmt
	getFollowersCountStmt         *sql.Stmt
	getFollowingStmt              *sql.Stmt
	getFollowingIdsStmt           *sql.Stmt
	getFollowingReactionStmt      *sql.Stmt
	getFollwersStmt               *sql.Stmt
	getMutualStmt                 *sql.Stmt
	getPollsStmt                  *sql.Stmt
	getPostsStmt                  *sql.Stmt
	getRatingStmt                 *sql.Stmt
	getUserCommentsStmt           *sql.Stmt
	getUserLikesStmt              *sql.Stmt
	getUsersStmt                  *sql.Stmt
	insertAdminStmt               *sql.Stmt
	insertTrendingStmt            *sql.Stmt
	insertTrendingUsersStmt       *sql.Stmt
	likeTrendingStmt              *sql.Stmt
	likeTrendingUsersStmt         *sql.Stmt
	updateRatingStmt              *sql.Stmt
	upsertPostRecommendationsStmt *sql.Stmt
	upsertUserRecommendationsStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		commentTrendingStmt:           q.commentTrendingStmt,
		commentTrendingUsersStmt:      q.commentTrendingUsersStmt,
		createPollsStmt:               q.createPollsStmt,
		deleteExpiredTokensStmt:       q.deleteExpiredTokensStmt,
		getAdminStmt:                  q.getAdminStmt,
		getFeedbackStmt:               q.getFeedbackStmt,
		getFollowersCountStmt:         q.getFollowersCountStmt,
		getFollowingStmt:              q.getFollowingStmt,
		getFollowingIdsStmt:           q.getFollowingIdsStmt,
		getFollowingReactionStmt:      q.getFollowingReactionStmt,
		getFollwersStmt:               q.getFollwersStmt,
		getMutualStmt:                 q.getMutualStmt,
		getPollsStmt:                  q.getPollsStmt,
		getPostsStmt:                  q.getPostsStmt,
		getRatingStmt:                 q.getRatingStmt,
		getUserCommentsStmt:           q.getUserCommentsStmt,
		getUserLikesStmt:              q.getUserLikesStmt,
		getUsersStmt:                  q.getUsersStmt,
		insertAdminStmt:               q.insertAdminStmt,
		insertTrendingStmt:            q.insertTrendingStmt,
		insertTrendingUsersStmt:       q.insertTrendingUsersStmt,
		likeTrendingStmt:              q.likeTrendingStmt,
		likeTrendingUsersStmt:         q.likeTrendingUsersStmt,
		updateRatingStmt:              q.updateRatingStmt,
		upsertPostRecommendationsStmt: q.upsertPostRecommendationsStmt,
		upsertUserRecommendationsStmt: q.upsertUserRecommendationsStmt,
	}
}
