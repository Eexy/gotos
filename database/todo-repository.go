package database

import "github.com/eexy/gotos/model"

type TodoRepository struct {
	Database *Database
}

func (r *TodoRepository) Add(title string) *model.Todo {
	todos := r.Database.LoadDb()
	todo := &model.Todo{Title: title, Completed: false}
	todos = append(todos, todo)
	r.Database.SaveDb(todos)
	return todo
}
