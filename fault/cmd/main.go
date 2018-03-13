package main

import (
	"fmt"

	service "github.com/burxtx/fault/fault/cmd/service"
	"github.com/burxtx/fault/fault/config"
	"github.com/burxtx/fault/fault/pkg/db"
)

func main() {
	config.Init("development")
	c := config.GetConfig()
	datasource := c.GetString("database.host")
	db, err := db.Init(datasource)
	if err != nil {
		panic(fmt.Errorf("Fatal error database connection: %s \n", err))
	}
	service.Run(db)
}
