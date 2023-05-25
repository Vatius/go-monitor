package internal

import (
	"CheckService/internal/entity"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	GetById(ctx context.Context, id string) (error, *entity.User)
	DeleteById(ctx context.Context, id string) error
}
