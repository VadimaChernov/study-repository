package scanner

import (
	"fmt"
	"todoapp/komandi"
)

func printPromt() {
	fmt.Println("введите команду: ")
}
func printExit() {
	fmt.Println("вы завершили программу")
}
func printAdd(title string) {
	fmt.Println("задача: " + title + " успешно добавлена ")
}
func printTasks(tasks map[string]komandi.Task) {
	fmt.Println("список задач: ", tasks)
}
func printDone(title string) {
	fmt.Println("задача: " + title + " помечена как выполненная")
}
func printDel(title string) {
	fmt.Println("задача: '" + title + "' успешно удалена")
}
func printHelp() {
	fmt.Println("exit")
	fmt.Println("help")
	fmt.Println("list")
	fmt.Println("done")
	fmt.Println("del")
	fmt.Println("add")
	fmt.Println("events")
}
func printEvents(events []Event) {
	fmt.Println("события", events)
}
