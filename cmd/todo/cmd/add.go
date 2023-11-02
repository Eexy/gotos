package cmd

import (
	"github.com/eexy/gotos/config"
	"github.com/spf13/cobra"
)

func NewAddTodoCmd(env *config.Env) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add todo",
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				env.Logger.Fatalln(err.Error())
			}

			todo := env.TodoRepository.Add(title)
			env.Logger.Println(todo)
		},
	}

	addCmd.Flags().StringP("title", "t", "", "Set new todo title")
	addCmd.MarkFlagRequired("title")

	return addCmd
}
