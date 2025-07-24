package scanner

import (
	"fmt"
	"os"

	"github.com/Onizukachi/GoToDo/tasks"
)

func printHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("help\t\t\t\tузнать доступные команды и их формат")
	fmt.Println("add {заголовок} {текст}\t\tдобавлять новые задачи в список задач")
	fmt.Println("list\t\t\t\tполучить полный список всех задач")
	fmt.Println("del {заголовок}\t\t\tудалить задачу по её заголовку")
	fmt.Println("events\t\t\t\tполучить список всех событий")
	fmt.Println("exit\t\t\t\tзавершить выполнение программы")
	fmt.Println()
}

func printGetInput() {
	fmt.Print("Введите команду: ")
}

func printTasks(tasks *[]tasks.Task) {
	fmt.Println("Cписок дел:")
	for i, task := range *tasks {
		fmt.Printf("%d. %s. %s. Статус: %v\n", i+1, task.Header, task.Body, task.IsDone)
	}
}

func printEmptyTaskList() {
	fmt.Println("Список задач пуст:(")
}

func printNotEnoughArgs() {
	fmt.Fprint(os.Stderr, "Не достаточно аргументов\n")
}

func printTaskAdded() {
	fmt.Println("Задача добавлена!")
}

func printTaskDeleted() {
	fmt.Println("Задача успешно удалена")
}

func printTaskDone() {
	fmt.Println("Задача выполнена")
}

func printEvents(events *[]Event) {
	fmt.Println("Список событий:")
	for _, event := range *events {
		fmt.Printf("%s\t%s\n", event.Time.Format("2006-01-02 15:04:05"), event.Description)
	}
}

func printEmptyEvents() {
	fmt.Println("Список событий пуст")
}

func printExit() {
	fmt.Println("Пока :) Буду ждать тебя снова!")
}

func printUnknownCmd() {
	fmt.Fprint(os.Stderr, "Такой команды я не знаю. Попробуй еще раз:(\n")
}

func printInputError(err *error) {
	fmt.Fprintf(os.Stderr, "ошибка чтения ввода %v", err)
}

func printTaskNotFound() {
	fmt.Println("Задача с таким заголовком не найдена :(")
}
