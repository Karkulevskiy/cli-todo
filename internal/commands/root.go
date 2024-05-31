package commands

import (
	"github.com/karkulevskiy/cli-todo/storage/sqlite"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo shows the list of tasks",
}

var session Session

type Session struct {
	Storage *sqlite.Storage
}

func InitSession(storage *sqlite.Storage) {
	session = Session{
		Storage: storage,
	}
}
