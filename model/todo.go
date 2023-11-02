package model

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
