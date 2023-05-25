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
	GetAll(ctx context.Context) (error, []entity.User)
}

type CheckListRepo interface {
	Create(ctx context.Context, service *entity.Alert) error
	Update(ctx context.Context, service *entity.Alert) error
	GetById(ctx context.Context, id string) (error, *entity.Alert)
	DeleteById(ctx context.Context, id string) error
	GetAll(ctx context.Context) (error, []entity.Alert)
	GetAllByUserID(ctx context.Context, userID string) (error, []entity.Alert)
}

type AlertRepo interface {
	Create(ctx context.Context, service *entity.Alert) error
	Update(ctx context.Context, service *entity.Alert) error
	GetById(ctx context.Context, id string) (error, *entity.Alert)
	DeleteById(ctx context.Context, id string) error
	GetAll(ctx context.Context) (error, []entity.Alert)
	GetAllByUserID(ctx context.Context, userID string) (error, []entity.Alert)
	GetAllByServiceID(ctx context.Context, serviceID string) (error, []entity.Alert)
}
