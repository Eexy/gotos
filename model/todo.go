package model

import "fmt"

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Priority  int    `json:"priority"`
	Completed bool   `json:"completed"`
}

func (t *Todo) String() string {
	return fmt.Sprintf("title: %s\tpriority: %d\tcompleted: %v", t.Title, t.Priority, t.Completed)
}
