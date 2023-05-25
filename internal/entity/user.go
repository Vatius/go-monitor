package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	UserName  string
	Password  string
	Token     string
	Email     string
	Telegram  string
	CreatedAt time.Time
	UpdateAt  time.Time
}
