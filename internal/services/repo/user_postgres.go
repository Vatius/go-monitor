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
	ErrorDuplicated       = errors.New("already exist")
	ErrorNotFound         = errors.New("not found")
	ErrorPass             = errors.New("wrong password")
	ErrorInvalidEmail     = errors.New("invalid email")
	ErrorInvalidOperation = errors.New("something go wrong")
	ErrorResetCode        = errors.New("wrong reset code")
	ErrorCantUpdate       = errors.New("can't update user")
	ErrorCantDelete       = errors.New("can't delete user")
	ErrorCantGet          = errors.New("can't get user")
	ErrorCantGetAll       = errors.New("can't get list of user")
)

const (
	userConflictErrMessage = "pq: duplicate key value violates unique constraint \"users_username_key\""
	userCreate             = "INSERT INTO users (id,username,password,email,telegramm) VALUES ($1, $2, $3, $4,$5);"
	userUpdate             = "UPDATE users SET username = $2, password = $3, email = $4, telegramm = $5 WHERE id = $1;"
	userDelete             = "DELETE FROM users WHERE id = $1 VALUES ($1);"
	userGet                = "SELECT id,username,password,token,email,telegram,created_at,updated_at FROM users WHERE id = $1;"
	userGetAll             = "SELECT id,username,password,token,email,telegram,created_at,updated_at FROM users;"
)

type (
	UserRepoImpl struct {
		db     *sql.DB
		logger logger.Interface
	}
)

func NewUserRepo(db *sql.DB, log logger.Interface) *UserRepoImpl {
	return &UserRepoImpl{
		db:     db,
		logger: log,
	}
}

func (s *UserRepoImpl) Create(ctx context.Context, user *entity.User) error {
	_, err := s.db.ExecContext(ctx, userCreate, user.ID, user.UserName, user.Password, user.Email, user.Telegram)
	if err != nil {
		if err.Error() == userConflictErrMessage {
			return ErrorDuplicated
		}
		return fmt.Errorf("can't creat user in db %v", err)
	}

	return nil
}

func (s *UserRepoImpl) Update(ctx context.Context, user *entity.User) error {
	_, err := s.db.ExecContext(ctx, userUpdate, user.ID, user.UserName, user.Password, user.Email, user.Telegram)
	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantUpdate, err)
	}

	return nil
}

func (s *UserRepoImpl) GetById(ctx context.Context, id string) (error, *entity.User) {
	res := s.db.QueryRowContext(ctx, userGet, id)
	u := entity.User{}
	err := res.Scan(&u.ID, &u.UserName, &u.Password, &u.Token, &u.Email, &u.Telegram, &u.CreatedAt, &u.UpdateAt)
	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGet, err), nil
	}

	return nil, &entity.User{
		ID:        u.ID,
		UserName:  u.UserName,
		Password:  u.Password,
		Token:     u.Token,
		Email:     u.Email,
		Telegram:  u.Telegram,
		CreatedAt: u.CreatedAt,
		UpdateAt:  u.UpdateAt,
	}
}

func (s *UserRepoImpl) DeleteById(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, userDelete, id)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantDelete, err)
	}

	return nil
}

func (s *UserRepoImpl) GetAll(ctx context.Context) (error, []entity.User) {
	userList, err := s.db.QueryContext(ctx, userGetAll)

	if err != nil {
		return fmt.Errorf("%s %v", ErrorCantGetAll, err), nil
	}
	u := entity.User{}
	res := []entity.User{}
	for userList.Next() {
		err = userList.Scan(&u.ID, &u.UserName, &u.Password, &u.Token, &u.Email, &u.Telegram, &u.CreatedAt, &u.UpdateAt)
		if err != nil {
			return fmt.Errorf("can't read storage response %v", err), nil
		}
		res = append(res, u)
	}

	return nil, res
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
