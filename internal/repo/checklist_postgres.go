package repo

import (
	"CheckService/internal/entity"
	"CheckService/pkg/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrorCantUpdateService = errors.New("can't update service")
	ErrorCantGetService    = errors.New("can't get service")
	ErrorCantDeleteService = errors.New("can't delete service")
	ErrorCantGetAllService = errors.New("can't get all services")
)

const (
	checkListConflictErrMessage = "pq: duplicate key value violates unique constraint \"users_username_key\""
	checkListCreate             = "INSERT INTO checklist (id,name,endpoint,description,status,icon,user_id) VALUES ($1, $2, $3, $4,$5,$6,$7);"
	checkListUpdate             = "UPDATE checklist SET name = $2, endpoint = $3, descrition = $4, status = $5 WHERE id = $1 ;"
	checkListDelete             = "DELETE FROM checklist WHERE id = $1 VALUES ($1)"
	checkListGet                = "SELECT id,name,endpoint,description,status,last_update,icon,user_id,created_at,updated_at FROM checklist WHERE id = $1 "
	checkListGetAll             = "SELECT id,name,endpoint,description,status,last_update,icon,user_id,created_at,updated_at FROM checklist  "
	checkListGetAllByUserID     = "SELECT id,name,endpoint,description,status,last_update,icon,user_id,created_at,updated_at FROM checklist WHERE user_id = $1 "
)

type (
	CheckListRepoImpl struct {
		db     *sql.DB
		logger logger.Interface
	}
)

func NewCheckListRepo(db *sql.DB, log logger.Interface) *CheckListRepoImpl {
	return &CheckListRepoImpl{
		db:     db,
		logger: log,
	}
}

func (s *CheckListRepoImpl) Create(ctx context.Context, service *entity.CheckList) error {
	_, err := s.db.ExecContext(ctx, checkListCreate, service.ID, service.Name, service.Endpoint, service.Description, service.Status, service.Icon, service.UserID)
	if err != nil {
		if err.Error() == checkListConflictErrMessage {
			return ErrorDuplicated
		}
		return fmt.Errorf("the service was not added to tracking %v", err)
	}

	return nil
}

func (s *CheckListRepoImpl) Update(ctx context.Context, service *entity.CheckList) error {
	_, err := s.db.ExecContext(ctx, checkListUpdate, service.ID, service.Name, service.Endpoint, service.Description, service.Status, service.Icon, service.UserID)
	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantUpdateService, err)
	}
	return nil
}

func (s *CheckListRepoImpl) GetById(ctx context.Context, id string) (error, *entity.CheckList) {
	res := s.db.QueryRowContext(ctx, checkListGet, id)
	c := entity.CheckList{}
	err := res.Scan(&c.ID, &c.Name, &c.Endpoint, &c.Description, &c.Status, &c.LastUpDate, &c.Icon, &c.UserID, &c.CreatedAt, &c.UpDatedAt)
	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetService, err), nil
	}

	return nil, &entity.CheckList{
		ID:          c.ID,
		Name:        c.Name,
		Endpoint:    c.Endpoint,
		Description: c.Description,
		Status:      c.Status,
		LastUpDate:  c.LastUpDate,
		Icon:        c.Icon,
		UserID:      c.UserID,
		CreatedAt:   c.CreatedAt,
		UpDatedAt:   c.UpDatedAt,
	}
}

func (s *CheckListRepoImpl) DeleteById(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, checkListDelete, id)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantDeleteService, err)
	}

	return nil
}

func (s *CheckListRepoImpl) GetAll(ctx context.Context) (error, []entity.CheckList) {
	serviceList, err := s.db.QueryContext(ctx, checkListGetAll)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetAllService, err), nil
	}
	c := entity.CheckList{}
	res := []entity.CheckList{}
	for serviceList.Next() {
		err = serviceList.Scan(&c.ID, &c.Name, &c.Endpoint, &c.Description, &c.Status, &c.LastUpDate, &c.Icon, &c.UserID, &c.CreatedAt, &c.UpDatedAt)
		if err != nil {
			return fmt.Errorf("can't read storage response %v", err), nil
		}
		res = append(res, c)
	}

	return nil, res
}

func (s *CheckListRepoImpl) GetAllByUserID(ctx context.Context, userID string) (error, []entity.CheckList) {
	serviceList, err := s.db.QueryContext(ctx, checkListGetAllByUserID, userID)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetAllService, err), nil
	}
	c := entity.CheckList{}
	res := []entity.CheckList{}
	for serviceList.Next() {
		err = serviceList.Scan(&c.ID, &c.Name, &c.Endpoint, &c.Description, &c.Status, &c.LastUpDate, &c.Icon, &c.UserID, &c.CreatedAt, &c.UpDatedAt)
		if err != nil {
			return fmt.Errorf("can't read storage response %v", err), nil
		}
		res = append(res, c)
	}

	return nil, res
}
