package main

import (
	"github.com/Onizukachi/GoToDo/scanner"
	"github.com/Onizukachi/GoToDo/tasks"
)

func main() {
	tasks := tasks.TaskList{}
	scanner := scanner.NewScanner(tasks)
	scanner.Run()
}
