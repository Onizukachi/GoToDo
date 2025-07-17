package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	Header    string
	Body      string
	CreatedAt time.Time
	IsDone    bool
	DoneAt    time.Time
}

func newTask(header, body string) *Task {
	return &Task{
		Header:    header,
		Body:      body,
		CreatedAt: time.Now(),
		IsDone:    false,
	}
}

type TaskList struct {
	Tasks []Task
}

func (taskList *TaskList) Add(task *Task) {
	taskList.Tasks = append(taskList.Tasks, *task)
}

type Event struct {
	Time time.Time
	Name string
}

func newEvent(name string) *Event {
	return &Event{
		Time: time.Now(),
		Name: name,
	}
}

type Menu struct {
	Events []Event
}

func (menu *Menu) LogEvent(event *Event) {
	menu.Events = append(menu.Events, *event)
}

func (menu Menu) GetUserInput() ([]string, error) {
	fmt.Print("Введите команду: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения ввода %v", err)
	}

	args := strings.Fields(scanner.Text())
	return args, nil
}

func main() {
	menu := Menu{}
	taskList := TaskList{}
	for {
		args, err := menu.GetUserInput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при чтении команды: %v\n", err)
			continue
		}

		event := newEvent(strings.Join(args, " "))
		menu.LogEvent(event)

		cmd := args[0]

		switch cmd {
		case "help":
			fmt.Println("Доступные команды:")
			fmt.Println("help\t\t\t\tузнать доступные команды и их формат")
			fmt.Println("add {заголовок} {текст}\t\tдобавлять новые задачи в список задач")
			fmt.Println("list\t\t\t\tполучить полный список всех задач")
			fmt.Println("del {заголовок}\t\t\tудалить задачу по её заголовку")
			fmt.Println("events\t\t\t\tполучить список всех событий")
			fmt.Println("exit\t\t\t\tзавершить выполнение программы")
			fmt.Println()
			continue
		case "list":
			if len(taskList.Tasks) == 0 {
				fmt.Println("Список задач пуст:(\n\n")
				continue
			}

			fmt.Println("Cписок дел:")
			for i, task := range taskList.Tasks {
				fmt.Printf("%d. %s. %s. Статус: %v\n", i, task.Header, task.Body, task.IsDone)
			}
			fmt.Println()
			continue
		case "add":
			if len(args[1:]) < 2 {
				fmt.Fprint(os.Stderr, "Не достаточно аргументов\n\n")
				continue
			}
			task := newTask(args[0], strings.Join(args[1:], " "))
			taskList.Add(task)
			fmt.Println("Задача добавлена!\n\n")
			continue
		case "del":
			fmt.Println("Удаляем дело:")
		case "done":
			fmt.Println("Завершаем дело:")
		case "events":
			if len(menu.Events) == 0 {
				fmt.Println("Список событий пуст")
				continue
			}

			fmt.Println("Список событий:")
			for i, event := range menu.Events {
				fmt.Printf("%d. %v\t%s\n", i, event.Time, event.Name)
			}
			fmt.Println()
			continue
		case "exit":
			fmt.Println("Пока :) Буду ждать тебя снова!")
			os.Exit(0)
		default:
			fmt.Fprint(os.Stderr, "Такой команды я не знаю. Попробуй еще раз:(\n\n")
			continue
		}
	}
}
