package internal

import (
	"CheckService/internal/entity"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, user *entity.User) error
	Update
	GetbyId
	Delete
}
