// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	null "gopkg.in/guregu/null.v4"
)

type NotificationsStatus string

const (
	NotificationsStatusUnseen NotificationsStatus = "Unseen"
	NotificationsStatusSeen   NotificationsStatus = "Seen"
	NotificationsStatusRead   NotificationsStatus = "Read"
)

func (e *NotificationsStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NotificationsStatus(s)
	case string:
		*e = NotificationsStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for NotificationsStatus: %T", src)
	}
	return nil
}

type AdminUser struct {
	MailID string `json:"mail_id"`
	OpenTo string `json:"open_to"`
}

type ChildComment struct {
	CommentID       int64     `json:"comment_id"`
	UserID          int32     `json:"user_id"`
	ParentCommentID int64     `json:"parent_comment_id"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
}

type DraftStory struct {
	DraftID   int32     `json:"draft_id"`
	MailID    string    `json:"mail_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Favourite struct {
	UserID int32 `json:"user_id"`
	PostID int64 `json:"post_id"`
}

type Feedback struct {
	FeedbackID int64       `json:"feedback_id"`
	UserID     int32       `json:"user_id"`
	ImageUri   null.String `json:"image_uri"`
	Content    null.String `json:"content"`
	CreatedAt  time.Time   `json:"created_at"`
	Status     bool        `json:"status"`
	Comment    null.String `json:"comment"`
}

type Like struct {
	UserID    int32     `json:"user_id"`
	PostID    int64     `json:"post_id"`
	Type      int32     `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Network struct {
	ID          int64     `json:"id"`
	FollowerID  int32     `json:"follower_id"`
	FollowingID int32     `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type Notification struct {
	ID        int64               `json:"id"`
	ForID     int32               `json:"for_id"`
	Status    NotificationsStatus `json:"status"`
	Metadata  json.RawMessage     `json:"metadata"`
	CreatedAt time.Time           `json:"created_at"`
}

type ParentComment struct {
	CommentID int64     `json:"comment_id"`
	UserID    int32     `json:"user_id"`
	PostID    int64     `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Poll struct {
	PollID       int32     `json:"poll_id"`
	PollBy       string    `json:"poll_by"`
	PollQuestion string    `json:"poll_question"`
	OptionsCount int32     `json:"options_count"`
	Options      string    `json:"options"`
	CreatedAt    time.Time `json:"created_at"`
}

type PollsReaction struct {
	PollID int32 `json:"poll_id"`
	UserID int32 `json:"user_id"`
	Type   int32 `json:"type"`
}

type Post struct {
	UserID     int32          `json:"user_id"`
	PostID     int64          `json:"post_id"`
	Content    sql.NullString `json:"content"`
	Media      sql.NullString `json:"media"`
	IsPrivate  bool           `json:"is_private"`
	OriginalID sql.NullInt64  `json:"original_id"`
	Edits      int32          `json:"edits"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type Postrecommendation struct {
	UserID              int32  `json:"user_id"`
	PostRecommendations string `json:"post_recommendations"`
}

type Profile struct {
	UserID          int32          `json:"user_id"`
	Username        string         `json:"username"`
	EmailID         string         `json:"email_id"`
	ProfileImageUri sql.NullString `json:"profile_image_uri"`
	Bio             sql.NullString `json:"bio"`
	CricIndex       int32          `json:"cric_index"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Following       int32          `json:"following"`
	Followers       int32          `json:"followers"`
	Interests       sql.NullString `json:"interests"`
}

type Story struct {
	StoryID   int32     `json:"story_id"`
	MailID    string    `json:"mail_id"`
	Content   string    `json:"content"`
	Media     string    `json:"media"`
	CreatedAt time.Time `json:"created_at"`
}

type Token struct {
	ScreenID  int64        `json:"screen_id"`
	UserID    int32        `json:"user_id"`
	TokenID   string       `json:"token_id"`
	CreatedAt time.Time    `json:"created_at"`
	ExpiredAt sql.NullTime `json:"expired_at"`
}

type TrendingPost struct {
	TpID      int64     `json:"tp_id"`
	PostID    int64     `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type TrendingUser struct {
	TuID      int64     `json:"tu_id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	UserID           int32          `json:"user_id"`
	MailID           string         `json:"mail_id"`
	Password         sql.NullString `json:"password"`
	CreatedAt        time.Time      `json:"created_at"`
	IdentityProvider sql.NullString `json:"identity_provider"`
	SubjectID        sql.NullString `json:"subject_id"`
}

type UserRecommendation struct {
	UserID    int32  `json:"user_id"`
	Recommend string `json:"recommend"`
}
