// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	ID         int64              `json:"id"`
	FromUserID int64              `json:"from_user_id"`
	ToUserID   int64              `json:"to_user_id"`
	IsSent     bool               `json:"is_sent"`
	Content    string             `json:"content"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	ID             int64            `json:"id"`
	Email          string           `json:"email"`
	Username       string           `json:"username"`
	Password       string           `json:"password"`
	UsersPhotoLink pgtype.Text      `json:"users_photo_link"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}