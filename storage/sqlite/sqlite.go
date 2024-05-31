package sqlite

import (
	"database/sql"

	"github.com/karkulevskiy/cli-todo/internal/domain"
	"github.com/karkulevskiy/cli-todo/storage"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(dbType string) (*Storage, error) {
	db, err := sql.Open("sqlite3", dbType)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks
		(
			id INTEGER PRIMARY KEY,
			task TEXT,
			time TIMESTAMP
		)`)

	if err != nil {
		db.Close()
		return nil, err
	}

	return &Storage{db}, nil
}

//TODO: update task

func (s *Storage) AddTask(task domain.Task) (int64, error) {
	res, err := s.db.Exec(`INSERT INTO tasks (task, time) VALUES (?, ?)`, task.Task, task.Time)
	if err != nil {
		return 0, err
	}

	taskID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return taskID, nil
}

func (s *Storage) DoneTask(id int64) error {
	res, err := s.db.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return storage.ErrTaskNotFound
	}

	return nil
}

func (s *Storage) Tasks() ([]domain.Task, error) {
	rows, err := s.db.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}

	var tasks []domain.Task

	var taskID int
	var taskText sql.NullString
	var taskTime sql.NullTime

	for rows.Next() {
		err := rows.Scan(&taskID, &taskText, &taskTime)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, domain.Task{
			ID:   taskID,
			Task: taskText.String,
			Time: taskTime.Time,
		})
	}

	return tasks, nil
}

func (s *Storage) RemoveAllTasks() error {
	_, err := s.db.Exec(`DROP TABLE tasks`)
	if err != nil {
		return err
	}

	return nil
}
