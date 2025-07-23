package tasks

import "time"

type Task struct {
	Header    string
	Body      string
	CreatedAt time.Time
	IsDone    bool
	DoneAt    time.Time
}

func NewTask(header, body string) *Task {
	return &Task{
		Header:    header,
		Body:      body,
		CreatedAt: time.Now(),
		IsDone:    false,
	}
}
