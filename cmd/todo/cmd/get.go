package cmd

import (
	"fmt"

	"github.com/eexy/gotos/config"
	"github.com/spf13/cobra"
)

func NewGetTodoCmd(env *config.Env) *cobra.Command {
	getTodoCmd := &cobra.Command{
		Use:   "get",
		Short: "Get todos",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flags().Changed("id") {
				id, err := cmd.Flags().GetInt("id")
				if err != nil {
					env.Logger.Fatalf(err.Error())
				}

				todo, err := env.TodoRepository.GetById(id)
				if err != nil {
					env.Logger.Fatal(err.Error())
				}

				env.Logger.Println(todo)
				return
			}

			todos := env.TodoRepository.Get()
			for _, todo := range todos {
				fmt.Println(todo)
			}
		},
	}

	getTodoCmd.Flags().Int("id", -1, "Todo's id to get")

	return getTodoCmd
}
