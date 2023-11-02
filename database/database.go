package database

import (
	"encoding/json"
	"log"
	"os"

	"github.com/eexy/gotos/model"
)

type Database struct {
	logger *log.Logger
}

func InitDb(logger *log.Logger) *Database {
	return &Database{logger: logger}
}

func (db *Database) LoadDb() []*model.Todo {
	todos := []*model.Todo{}
	fileContent, err := os.ReadFile("gotos.json")
	if err != nil {
		db.logger.Printf("Unable to read gotos.json file. Err : %s", err.Error())
		return todos
	}

	err = json.Unmarshal(fileContent, &todos)
	if err != nil {
		db.logger.Printf("Invalid json in gotos.json. Err :%s\n", err.Error())
	}

	return todos
}

func (db *Database) SaveDb(todos []*model.Todo) {
	json, err := json.Marshal(todos)

	if err != nil {
		db.logger.Fatalf("Unable to convert data to json. Err :%s\n", err.Error())
	}

	err = os.WriteFile("gotos.json", json, 0666)
	if err != nil {
		db.logger.Fatalf("Unable to save todos. Err %s", err)
	}

	db.logger.Println("Todos saved")
}
