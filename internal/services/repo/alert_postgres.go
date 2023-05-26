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
	ErrorCantUpdateAlert = errors.New("can't update alert")
	ErrorCantGetAlert    = errors.New("can't get alert")
	ErrorCantDeleteAlert = errors.New("can't delete alert")
	ErrorCantGetAllAlert = errors.New("can't get all alert")
)

const (
	alertConflictErrMessage = "pq: duplicate key value violates unique constraint \"users_username_key\""
	alertCreate             = "INSERT INTO alert (id,check_id,name,description,status) VALUES ($1, $2, $3, $4,$5,$6,$7);"
	alertUpdate             = "UPDATE alert SET name = $2, description = $3, status = $4 WHERE id = $1;"
	alertDelete             = "DELETE FROM alert WHERE id = $1 VALUES ($1);"
	alertGet                = "SELECT id,check_id,alert_date,name,description,status,created_at FROM alert WHERE id = $1;"
	alertGetAll             = "SELECT id,check_id,alert_date,name,description,status,created_at FROM alert; "
	alertGetAllByUserID     = "SELECT id,check_id,alert_date,name,description,status,created_at FROM alert INNER JOIN checklist on alert.check_id=checklist.id AND WHERE user_id = $1;"
	alertGetAllByServiceID  = "SELECT id,check_id,alert_date,name,description,status,created_at FROM alert WHERE check_id = $1;"
)

type AlertRepoImpl struct {
	db     *sql.DB
	logger logger.Interface
}

func NewAlertRepo(db *sql.DB, log logger.Interface) *AlertRepoImpl {
	return &AlertRepoImpl{
		db:     db,
		logger: log,
	}
}

func (s *AlertRepoImpl) Create(ctx context.Context, alert *entity.Alert) error {
	_, err := s.db.ExecContext(ctx, alertCreate, alert.ID, alert.CheckID, alert.AlertDate, alert.Name, alert.Description, alert.Status, alert.CreatedAt)
	if err != nil {
		if err.Error() == checkListConflictErrMessage {
			return ErrorDuplicated
		}
		return fmt.Errorf("the alert din't create %v", err)
	}

	return nil
}

func (s *AlertRepoImpl) Update(ctx context.Context, alert *entity.Alert) error {
	_, err := s.db.ExecContext(ctx, alertUpdate, alert.Name, alert.Description, alert.Status)
	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantUpdateAlert, err)
	}
	return nil
}

func (s *AlertRepoImpl) GetById(ctx context.Context, id string) (error, *entity.Alert) {
	res := s.db.QueryRowContext(ctx, alertGet, id)
	a := entity.Alert{}
	err := res.Scan(&a.ID, &a.CheckID, &a.AlertDate, &a.Name, &a.Description, &a.Status, &a.CreatedAt)
	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetAlert, err), nil
	}

	return nil, &entity.Alert{
		ID:          a.ID,
		CheckID:     a.CheckID,
		AlertDate:   a.AlertDate,
		Name:        a.Name,
		Description: a.Description,
		Status:      a.Status,
		CreatedAt:   a.CreatedAt,
	}

}

func (s *AlertRepoImpl) DeleteById(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, alertDelete, id)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantDeleteAlert, err)
	}

	return nil
}

func (s *AlertRepoImpl) GetAll(ctx context.Context) (error, []entity.Alert) {
	alertList, err := s.db.QueryContext(ctx, alertGetAll)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetAllAlert, err), nil
	}
	a := entity.Alert{}
	res := []entity.Alert{}
	for alertList.Next() {
		err = alertList.Scan(&a.ID, &a.CheckID, &a.AlertDate, &a.Name, &a.Description, &a.Status, &a.CreatedAt)
		if err != nil {
			return fmt.Errorf("can't read storage response %v", err), nil
		}
		res = append(res, a)
	}

	return nil, res

}

func (s *AlertRepoImpl) GetAllByUserID(ctx context.Context, userID string) (error, []entity.Alert) {
	alertList, err := s.db.QueryContext(ctx, alertGetAllByUserID, userID)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetAllAlert, err), nil
	}
	a := entity.Alert{}
	res := []entity.Alert{}
	for alertList.Next() {
		err = alertList.Scan(&a.ID, &a.CheckID, &a.AlertDate, &a.Name, &a.Description, &a.Status, &a.CreatedAt)
		if err != nil {
			return fmt.Errorf("can't read storage response %v", err), nil
		}
		res = append(res, a)
	}

	return nil, res

}

func (s *AlertRepoImpl) GetAllByServiceID(ctx context.Context, serviceID string) (error, []entity.Alert) {
	alertList, err := s.db.QueryContext(ctx, alertGetAllByServiceID, serviceID)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetAllAlert, err), nil
	}
	a := entity.Alert{}
	res := []entity.Alert{}
	for alertList.Next() {
		err = alertList.Scan(&a.ID, &a.CheckID, &a.AlertDate, &a.Name, &a.Description, &a.Status, &a.CreatedAt)
		if err != nil {
			return fmt.Errorf("can't read storage response %v", err), nil
		}
		res = append(res, a)
	}

	return nil, res

}
