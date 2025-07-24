package tasks

import "time"

type Task struct {
	Header    string
	Body      string
	CreatedAt time.Time
	IsDone    bool
	DoneAt    time.Time
}

func (t *Task) MarkDone() {
	t.IsDone = true
	t.DoneAt = time.Now()
}

func (t *Task) UnmarkDone() {
	t.IsDone = false
	t.DoneAt = time.Time{}
}

func NewTask(header, body string) *Task {
	return &Task{
		Header:    header,
		Body:      body,
		CreatedAt: time.Now(),
		IsDone:    false,
	}
}
