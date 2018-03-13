package main

import (
	service "github.com/burxtx/fault/todo/cmd/service"
	"github.com/burxtx/fault/todo/config"
	"github.com/burxtx/fault/todo/pkg/db"
)

func main() {
	config.Init("development")
	c := config.GetConfig()
	datasource := c.GetString("database.host")
	db, err := db.Init(datasource)
	service.Run(db)
}
