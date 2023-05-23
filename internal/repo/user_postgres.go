package repo

import (
	"CheckService/internal/entity"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/mail"
)

var (
	ErrorDuplicated       = errors.New("already exist")
	ErrorNotFound         = errors.New("not found")
	ErrorPass             = errors.New("wrong password")
	ErrorInvalidEmail     = errors.New("invalid email")
	ErrorInvalidOperation = errors.New("something go wrong")
	ErrorResetCode        = errors.New("wrong reset code")
)

const (
	conflictErrMessage = "pq: duplicate key value violates unique constraint \"users_username_key\""
)

type (
	Service struct {
		db        *sql.DB
		key       string
		tableName string
	}
)

func NewService(db *sql.DB, tableName, key string) *Service {
	return &Service{
		db:        db,
		key:       key,
		tableName: tableName,
	}
}

func (s *Service) Create(username, pass, email, telegram string) (*entity.User, error) {
	if !validateEmail(email) {
		return nil, ErrorInvalidEmail
	}
	id := uuid.New()
	hashedPass := hashPass(s.key, pass)
	_, err := s.db.Exec(fmt.Sprintf("INSERT INTO %s (id,username,password,email,telegramm) VALUES ($1, $2, $3, $4,$5);", s.tableName), id.String(), username, hashedPass, email, telegram)
	if err != nil {
		if err.Error() == conflictErrMessage {
			return nil, ErrorDuplicated
		}
		return nil, fmt.Errorf("can't creat user in db %v", err)
	}

	return &entity.User{
		ID:       id,
		UserName: username,
		Email:    email,
		Telegram: telegram,
	}, nil
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPass(key, pass string) string {
	hash := md5.Sum([]byte(key + pass))
	return hex.EncodeToString(hash[:])
}
