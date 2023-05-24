package entity

import (
	"github.com/google/uuid"
	"time"
)

type CheckList struct {
	ID          uuid.UUID
	Name        string
	Endpoint    string
	Description string
	Status      string
	LastUpDate  time.Time
	Icon        string
	UserID      uuid.UUID
	CreatedAt   time.Time
	UpDatedAt   time.Time
}
