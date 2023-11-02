package main

import (
	"log"

	"github.com/eexy/gotos/cmd"
	"github.com/eexy/gotos/config"
	"github.com/eexy/gotos/database"
)

func main() {
	// Set app logger
	logger := config.InitAppLogger("logs.txt")

	// Set app environment
	db := database.InitDb(logger)
	todoRepository := &database.TodoRepository{Database: db}
	appEnv := &config.Env{Logger: logger, TodoRepository: todoRepository}

	root := cmd.NewRootCmd(appEnv)
	if err := root.Execute(); err != nil {
		log.Fatalln(err)
	}
}
