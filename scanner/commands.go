package scanner

import (
	"fmt"
	"os"
	"strings"

	"github.com/Onizukachi/GoToDo/tasks"
)

func (s *Scanner) cmdHelp() {
	printHelp()
}

func (s *Scanner) cmdList() {
	if len(s.taskList.Tasks) == 0 {
		printEmptyTaskList()
		return
	}

	printTasks(s.taskList.Tasks)
}

func (s *Scanner) cmdAdd(args []string) {
	if len(args) < 3 {
		printNotEnoughArgs()
		return
	}

	newTask := tasks.NewTask(args[1], strings.Join(args[2:], " "))
	err := s.taskList.Add(newTask)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка при добавлении задачи: %v\n", err)
		return
	}

	printTaskAdded()
}

func (s *Scanner) cmdDel(args []string) {
	if len(args) < 2 {
		printNotEnoughArgs()
		return
	}

	err := s.taskList.Del(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка при удалении задачи: %v\n", err)
		return
	}

	printTaskDeleted()
}

func (s *Scanner) cmdMarkDone(args []string) {
	if len(args) < 2 {
		printNotEnoughArgs()
		return
	}

	task, ok := s.taskList.Tasks[args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, "Задача не найдена\n")
		return
	}

	task.MarkDone()
	printTaskDone()
}

func (s *Scanner) listEvents() {
	if len(s.events) == 0 {
		printEmptyEvents()
	} else {
		printEvents(&s.events)
	}
}
