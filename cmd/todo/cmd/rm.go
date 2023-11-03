package cmd

import (
	"github.com/eexy/gotos/config"
	"github.com/spf13/cobra"
)

func NewRemoveTodoCmd(env *config.Env) *cobra.Command {
	removeTodoCmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove todo by id and return it",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetInt("id")
			if err != nil {
				env.Logger.Fatalln(err.Error())
			}

			todo, err := env.TodoRepository.Remove(id)
			if err != nil {
				env.Logger.Fatalf(err.Error())
			}

			env.Logger.Printf("Todo %d has been deleted\n", todo.Id)
		},
	}

	removeTodoCmd.Flags().Int("id", -1, "Todo's id to delete")
	removeTodoCmd.MarkFlagRequired("id")

	return removeTodoCmd
}
