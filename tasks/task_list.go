package tasks

import (
	"errors"
)

type TaskList struct {
	Tasks map[string]*Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: make(map[string]*Task),
	}
}

func (list *TaskList) Add(newTask *Task) error {
	if _, ok := list.Tasks[newTask.Header]; ok {
		return errors.New("задача с таким заголовком уже существует")
	}

	list.Tasks[newTask.Header] = newTask
	return nil
}

func (list *TaskList) Del(title string) error {
	if _, ok := list.Tasks[title]; !ok {
		return errors.New("задачи с таким заголовком не существует")
	}

	delete(list.Tasks, title)
	return nil
}
