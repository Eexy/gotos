package cmd

import (
	"fmt"

	"github.com/eexy/gotos/cmd/todo"
	"github.com/eexy/gotos/config"
	"github.com/spf13/cobra"
)

func NewRootCmd(env *config.Env) *cobra.Command {
	root := &cobra.Command{
		Use:   "gotos",
		Short: "Handle your todos from cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from root")
		},
	}

	root.AddCommand(todo.NewTodoCmd(env))
	root.PersistentFlags().String("file", "./gotos.json", "Path to JSON file to use as DB")
	return root
}
