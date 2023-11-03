package database

import (
	"errors"
	"fmt"

	"github.com/eexy/gotos/model"
)

type TodoRepository struct {
	Database *Database
}

func (r *TodoRepository) Add(title string) *model.Todo {
	todos := r.Database.LoadDb()
	todo := &model.Todo{Id: r.generateNewTodoId(todos), Title: title, Completed: false}
	todos = append(todos, todo)
	r.Database.SaveDb(todos)
	return todo
}

func (r *TodoRepository) generateNewTodoId(todos []*model.Todo) int {
	if len(todos) == 0 {
		return 0
	}

	max := todos[0].Id
	for _, todo := range todos {
		if todo.Id > max {
			max = todo.Id
		}
	}

	return max + 1
}

func (r *TodoRepository) Get() []*model.Todo {
	return r.Database.LoadDb()
}

func (r *TodoRepository) GetById(id int) (*model.Todo, error) {
	todos := r.Database.LoadDb()

	for _, todo := range todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Unable to find todo with id %d\n", id))
}

func (r *TodoRepository) Remove(id int) (*model.Todo, error) {
	todos := r.Database.LoadDb()
	filteredTodos := []*model.Todo{}
	var deletedTodo *model.Todo = nil
	for _, todo := range todos {
		if todo.Id == id {
			deletedTodo = todo
		} else {
			filteredTodos = append(filteredTodos, todo)
		}
	}

	if deletedTodo == nil {
		return nil, errors.New(fmt.Sprintf("Unable to find todo with id %d to delete\n", id))
	}

	r.Database.SaveDb(filteredTodos)
	return deletedTodo, nil
}
