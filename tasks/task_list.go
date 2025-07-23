package tasks

type TaskList struct {
	Tasks []Task
}

func (taskList *TaskList) Add(task *Task) {
	taskList.Tasks = append(taskList.Tasks, *task)
}
