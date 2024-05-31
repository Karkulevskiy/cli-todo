package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/karkulevskiy/cli-todo/internal/domain"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add adds a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		if task == "" {
			fmt.Println("Task can't be empty!")
			return
		}

		_, err := session.Storage.AddTask(domain.Task{Task: task, Time: time.Now()})
		if err != nil {
			fmt.Println("Internal error: ", err)
			return
		}

		fmt.Printf("Task: \"%s\" added!\n", task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
