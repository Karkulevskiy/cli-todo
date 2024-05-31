package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows the list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := session.Storage.Tasks()
		if err != nil {
			fmt.Println("internal error: ", err)
			return
		}
		fmt.Println("Your tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s - %v\n",
				task.ID,
				task.Task,
				task.Time.Format("2006-01-02 15:04:05"))
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
