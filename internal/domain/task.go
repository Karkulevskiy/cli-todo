package domain

import "time"

type Task struct {
	ID   int
	Task string
	Time time.Time
}
