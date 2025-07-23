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
	cmd := args[0]

	switch cmd {
	case "help":
		printHelp()
	case "list":
		if len(s.taskList.Tasks) == 0 {
			printEmptyTaskList()
		} else {
			printTasks(&s.taskList.Tasks)
		}
	case "add":
		if len(args[1:]) < 2 {
			printNotEnoughArgs()
		} else {
			newTask := tasks.NewTask(args[1], strings.Join(args[2:], " "))
			s.taskList.Add(newTask)
			printTaskAdded()
		}
	// TODO: del done
	case "del":
		printTaskDeleted()
	case "done":
		printTaskDone()
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
