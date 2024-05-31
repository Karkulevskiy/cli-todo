package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		err := session.Storage.RemoveAllTasks()
		if err != nil {
			fmt.Println("internal error: ", err)
			return
		}

		fmt.Println("All tasks deleted!")
	},
}

func init() {
	RootCmd.AddCommand(ClearCmd)
}
