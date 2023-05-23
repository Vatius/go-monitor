package main

import (
	"CheckService/storage"
	"encoding/json"
	"log"
	"os"
)

type (
	Config struct {
		DB                storage.ConnectionConfig `json:"db"`
		Key               string                   `json:"key"`
		UsersTableName    string                   `json:"users_table_name"`
		TasksTableName    string                   `json:"tasks_table_name"`
		CommentsTableName string                   `json:"comments_table_name"`
		HttpPort          int                      `json:"http_port"`
	}
)

func main() {
	log.Print("Start")
	config, err := GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := storage.Connect(config.DB)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	func (p *Service) CreateUser(req UserInfo) (*users.User, error) {
		return p.usersSVC.Create(req.UserName, req.UserPass)
	}
}

func GetConfig() (Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return Config{}, err
	}

	c := Config{}
	if err = json.Unmarshal(file, &c); err != nil {
		return Config{}, err
	}

	return c, nil
}
