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

func NewScanner(tasks *tasks.TaskList) *Scanner {
	return &Scanner{
		taskList:  *tasks,
		ioScanner: bufio.NewScanner(os.Stdin),
		events:    []Event{},
	}
}

func (s *Scanner) logEvent(event *Event) {
	s.events = append(s.events, *event)
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

		input := s.ioScanner.Text()
		s.logEvent(NewEvent(input))
		s.process(input)
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
		s.cmdHelp()
	case "list":
		s.cmdList()
	case "add":
		s.cmdAdd(args)
	case "del":
		s.cmdDel(args)
	case "done":
		s.cmdMarkDone(args)
	case "events":
		s.listEvents()
	case "exit":
		printExit()
		os.Exit(0)
	default:
		printUnknownCmd()
	}
}
