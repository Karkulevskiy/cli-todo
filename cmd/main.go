package main

import (
	cmds "github.com/karkulevskiy/cli-todo/internal/commands"
	"github.com/karkulevskiy/cli-todo/storage/sqlite"
)

// test
const (
	SQLite   = "tasks.db"
	InMemory = ":memory:"
)

func main() {
	cmds.InitSession(must(SQLite))
	cmds.RootCmd.Execute()
}

func must(dbType string) *sqlite.Storage {
	db, err := sqlite.New(dbType)
	if err != nil {
		panic(err)
	}

	return db
}
