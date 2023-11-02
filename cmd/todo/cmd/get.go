package cmd

import (
	"fmt"

	"github.com/eexy/gotos/config"
	"github.com/spf13/cobra"
)

func NewGetTodoCmd(env *config.Env) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get todos",
		Run: func(cmd *cobra.Command, args []string) {
			todos := env.TodoRepository.Get()

			for _, todo := range todos {
				fmt.Println(todo)
			}
		},
	}
}
