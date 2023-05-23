package repo

import (
	"CheckService/internal/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

var (
	ErrorDuplicated       = errors.New("already exist")
	ErrorNotFound         = errors.New("not found")
	ErrorPass             = errors.New("wrong password")
	ErrorInvalidEmail     = errors.New("invalid email")
	ErrorInvalidOperation = errors.New("something go wrong")
	ErrorResetCode        = errors.New("wrong reset code")
	ErrorCantUpdate       = errors.New("can't update user")
	ErrorCantDelete       = errors.New("can't delete user")
	ErrorCantGet          = errors.New("can't get user")
)

const (
	conflictErrMessage = "pq: duplicate key value violates unique constraint \"users_username_key\""
	create             = "INSERT INTO users (id,username,password,email,telegramm) VALUES ($1, $2, $3, $4,$5);"
	update             = "UPDATE users SET username = $2, password = $3, email = $4, telegramm = $5 WHERE id = $1 ;"
	delete             = "DELETE FROM users WHERE id = $1 VALUES ($1)"
	get                = "SELECT id,username,password,token,email,telegram,created_at,updated_at FROM users WHERE id = $1 "
)

type (
	UserRepoImpl struct {
		db     *sql.DB
		logger log.Logger
	}
)

func NewService(db *sql.DB, logger log.Logger) *UserRepoImpl {
	return &UserRepoImpl{
		db:     db,
		logger: logger,
	}
}

func (s *UserRepoImpl) Create(ctx context.Context, user *entity.User) error {
	_, err := s.db.Exec(create, user.ID, user.UserName, user.Password, user.Email, user.Telegram)
	if err != nil {
		if err.Error() == conflictErrMessage {
			return ErrorDuplicated
		}
		return fmt.Errorf("can't creat user in db %v", err)
	}

	return nil
}

func (s *UserRepoImpl) Update(ctx context.Context, user *entity.User) error {
	_, err := s.db.Exec(update, user.ID, user.UserName, user.Password, user.Email, user.Telegram)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantUpdate, err)
	}

	return nil
}

func (s *UserRepoImpl) GetById(ctx context.Context, id string) (error, *entity.User) {
	_, err := s.db.Query(get, id)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGet, err), nil
	}

	return nil, &entity.User{
		ID:        ,
		UserName:  "",
		Password:  "",
		Token:     "",
		Email:     "",
		Telegram:  "",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	}
}

func (s *UserRepoImpl) DeleteById(ctx context.Context, id string) error {
	_, err := s.db.Exec(delete, id)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantDelete, err)
	}

	return nil
}

//func validateEmail(email string) bool {
//	_, err := mail.ParseAddress(email)
//	return err == nil
//}

//func hashPass(key, pass string) string {
//	hash := md5.Sum([]byte(key + pass))
//	return hex.EncodeToString(hash[:])
//}
//hashedPass := hashPass(s.key, pass)
//id := uuid.New()
