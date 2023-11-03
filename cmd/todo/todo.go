package todo

import (
	"github.com/eexy/gotos/cmd/todo/cmd"
	"github.com/eexy/gotos/config"
	"github.com/spf13/cobra"
)

func NewTodoCmd(env *config.Env) *cobra.Command {
	todoCmd := &cobra.Command{
		Use:   "todos",
		Short: "Handle your todos",
		Run: func(cmd *cobra.Command, args []string) {
			env.Logger.Println("Hello from todos")
		},
	}

	todoCmd.AddCommand(cmd.NewAddTodoCmd(env))
	todoCmd.AddCommand(cmd.NewGetTodoCmd(env))
	todoCmd.AddCommand(cmd.NewRemoveTodoCmd(env))
	todoCmd.AddCommand(cmd.NewCompleteTodoCmd(env))

	return todoCmd
}
