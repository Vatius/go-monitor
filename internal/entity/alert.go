package entity

import (
	"github.com/google/uuid"
	"time"
)

type Alert struct {
	ID          uuid.UUID
	CheckID     uuid.UUID
	Date        time.Time
	Name        string
	Description string
	Status      string
	CreatedAt   time.Time
}
