package commands

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/karkulevskiy/cli-todo/storage"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "done",
	Short: "Marks the task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		tasksDone, err := validateArgs(args)
		if err != nil {
			return
		}

		notFoundTasks := strings.Builder{}
		foundTasks := strings.Builder{}

		for _, taskID := range tasksDone {
			err = session.Storage.DoneTask(taskID)
			if err != nil {
				if errors.Is(err, storage.ErrTaskNotFound) {
					notFoundTasks.WriteString(" " + strconv.Itoa(int(taskID)))
					continue
				}

				fmt.Println("internal error: ", err)
				return
			}
			foundTasks.WriteString(" " + strconv.Itoa(int(taskID)))
		}

		if foundTasks.String() != "" {
			fmt.Printf("Tasks: \"%s\" done!\n", foundTasks.String()[1:])
		}

		if notFoundTasks.String() != "" {
			fmt.Printf("Not found tasks: \"%s\"\n", notFoundTasks.String()[1:])
		}
	},
}

func validateArgs(args []string) ([]int64, error) {
	tasksDone := mapset.NewSet[int64]()

	for _, id := range args {
		taskId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("failed to parse task id: ", id)
			return nil, err
		}

		tasksDone.Add(taskId)
	}

	tasksDoneSlice := tasksDone.Clone().ToSlice()

	if len(tasksDoneSlice) == 0 {
		fmt.Println("No tasks to delete")
		return nil, errors.New("no tasks to delete")
	}

	return tasksDoneSlice, nil
}

func init() {
	RootCmd.AddCommand(doCmd)
}
