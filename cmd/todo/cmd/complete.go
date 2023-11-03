package cmd

import (
	"github.com/eexy/gotos/config"
	"github.com/spf13/cobra"
)

func NewCompleteTodoCmd(env *config.Env) *cobra.Command {
	completeTododCmd := &cobra.Command{
		Use:   "complete",
		Short: "Set the specify todo to completed",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetInt("id")
			if err != nil {
				env.Logger.Fatalln(err.Error())
			}

			todo, err := env.TodoRepository.Update(id, true)
			if err != nil {
				env.Logger.Fatalln(err)
			}

			env.Logger.Println(todo)
		},
	}

	completeTododCmd.Flags().Int("id", -1, "Todo's id to update")
	completeTododCmd.MarkFlagRequired("id")

	return completeTododCmd
}
