package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/eexy/gotos/cmd"
	"github.com/eexy/gotos/config"
	"github.com/eexy/gotos/database"
)

func main() {

	// Set app logger
	logger := config.InitAppLogger("logs.txt")

	// Set json file to use
	file := flag.String("file", "./gotos.json", "")
	flag.Parse()
	cwd, err := os.Getwd()
	if err != nil {
		logger.Println("Unable to get current working directory")
	}
	filePath := path.Join(cwd, *file)

	// Set app environment
	db := database.InitDb(filePath, logger)
	todoRepository := &database.TodoRepository{Database: db}
	appEnv := &config.Env{Logger: logger, TodoRepository: todoRepository}

	root := cmd.NewRootCmd(appEnv)
	if err := root.Execute(); err != nil {
		log.Fatalln(err)
	}
}
