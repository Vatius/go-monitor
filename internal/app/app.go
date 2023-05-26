package app

import (
	"CheckService/config"
	"CheckService/pkg/logger"
	"database/sql"
	"fmt"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	l.Info("starting app")
	db, err := Connect(cfg)
	if err != nil {
		l.Fatal("couldn't connect to DB %v", err)
	}
	defer db.Close()
}

func Connect(config *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PG.Host, config.PG.Port, config.PG.User, config.PG.Password, config.PG.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
