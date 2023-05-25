package app

import (
	"CheckService/config"
	"CheckService/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	l.Info("hello world")
	l.Error("произошла хуйня")
	l.Debug("фикс")
	l.Warn("жопа")

}
