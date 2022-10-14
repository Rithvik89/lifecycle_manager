// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
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

type NullNotificationsStatus struct {
	NotificationsStatus NotificationsStatus
	Valid               bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNotificationsStatus) Scan(value interface{}) error {
	if value == nil {
		ns.NotificationsStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NotificationsStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNotificationsStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.NotificationsStatus, nil
}

type ChildComment struct {
	CommentID       int64     `json:"comment_id"`
	UserID          int32     `json:"user_id"`
	ParentCommentID int64     `json:"parent_comment_id"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
}

type Like struct {
	UserID    int32     `json:"user_id"`
	PostID    int64     `json:"post_id"`
	Type      int32     `json:"type"`
	CreatedAt time.Time `json:"created_at"`
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

type Post struct {
	UserID     int32          `json:"user_id"`
	PostID     int64          `json:"post_id"`
	Content    sql.NullString `json:"content"`
	Media      sql.NullString `json:"media"`
	IsPrivate  bool           `json:"is_private"`
	OriginalID sql.NullInt64  `json:"original_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type Profile struct {
	UserID          int32          `json:"user_id"`
	Username        string         `json:"username"`
	EmailID         string         `json:"email_id"`
	ProfileImageUri sql.NullString `json:"profile_image_uri"`
	Bio             sql.NullString `json:"bio"`
	CricIndex       int32          `json:"cric_index"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Interests       sql.NullString `json:"interests"`
}

type Token struct {
	ScreenID  int64        `json:"screen_id"`
	UserID    int32        `json:"user_id"`
	TokenID   string       `json:"token_id"`
	CreatedAt time.Time    `json:"created_at"`
	ExpiredAt sql.NullTime `json:"expired_at"`
}

type User struct {
	UserID           int32          `json:"user_id"`
	MailID           string         `json:"mail_id"`
	Password         sql.NullString `json:"password"`
	CreatedAt        time.Time      `json:"created_at"`
	IdentityProvider sql.NullString `json:"identity_provider"`
	SubjectID        sql.NullString `json:"subject_id"`
}