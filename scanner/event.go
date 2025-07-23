package scanner

import "time"

type Event struct {
	Time        time.Time
	Description string
}

func NewEvent(description string) *Event {
	return &Event{
		Time:        time.Now(),
		Description: description,
	}
}
