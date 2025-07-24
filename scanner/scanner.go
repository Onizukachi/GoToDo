package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Onizukachi/GoToDo/tasks"
)

type Scanner struct {
	taskList  tasks.TaskList
	ioScanner *bufio.Scanner
	events    []Event
}

func (s *Scanner) logEvent(event *Event) {
	s.events = append(s.events, *event)
}

func NewScanner(tasks tasks.TaskList) *Scanner {
	return &Scanner{
		taskList:  tasks,
		ioScanner: bufio.NewScanner(os.Stdin),
		events:    []Event{},
	}
}

func (s *Scanner) Run() {
	for {
		printGetInput()
		s.ioScanner.Scan()

		if err := s.ioScanner.Err(); err != nil {
			s.logEvent(NewEvent(err.Error()))
			printInputError(&err)
			continue
		}

		s.logEvent(NewEvent(""))
		s.process(s.ioScanner.Text())
		fmt.Println()
	}
}

func (s *Scanner) process(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		printNotEnoughArgs()
		return
	}

	cmd := args[0]

	switch cmd {
	case "help":
		printHelp()
	case "list":
		if len(s.taskList.Tasks) == 0 {
			printEmptyTaskList()
			break
		}

		printTasks(&s.taskList.Tasks)
	case "add":
		if len(args) < 3 {
			printNotEnoughArgs()
			break
		}

		newTask := tasks.NewTask(args[1], strings.Join(args[2:], " "))
		s.taskList.Add(newTask)
		printTaskAdded()
	case "del":
		if len(args) < 2 {
			printNotEnoughArgs()
			break
		}

		found := false
		for i, task := range s.taskList.Tasks {
			if task.Header == args[1] {
				s.taskList.Tasks = append(s.taskList.Tasks[:i], s.taskList.Tasks[i+1:]...)
				printTaskDeleted()
				found = true
				break
			}
		}

		if !found {
			printTaskNotFound()
		}
	case "done":
		if len(args) < 2 {
			printNotEnoughArgs()
			break
		}

		found := false
		for i := range s.taskList.Tasks {
			if s.taskList.Tasks[i].Header == args[1] {
				s.taskList.Tasks[i].MarkDone()
				printTaskDone()
				found = true
				break
			}
		}

		if !found {
			printTaskNotFound()
		}
	case "events":
		if len(s.events) == 0 {
			printEmptyEvents()
		} else {
			printEvents(&s.events)
		}
	case "exit":
		printExit()
		os.Exit(0)
	default:
		printUnknownCmd()
	}
}
