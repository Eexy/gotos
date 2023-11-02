package config

import (
	"log"

	"github.com/eexy/gotos/database"
)

type Env struct {
	Logger         *log.Logger
	TodoRepository *database.TodoRepository
}
