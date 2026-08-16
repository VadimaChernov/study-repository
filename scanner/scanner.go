package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todoapp/komandi"
)

type Scanner struct {
	todoList *komandi.Komand
	events   []Event
}

func NewScanner(todoList *komandi.Komand) *Scanner {
	return &Scanner{
		todoList: todoList,
	}
}

func (s *Scanner) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPromt()
		ok := scanner.Scan()
		if !ok {
			break
		}
		inputString := scanner.Text()

		result := s.process(inputString)
		if result != "" {
			if result == needExit {
				printExit()
				return
			}
		}
		event := NewEvent(result, inputString)
		s.events = append(s.events, event)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("ошибка чтения", err)
	}

}

func (s *Scanner) process(inputString string) string {
	fields := strings.Fields(inputString)
	if len(fields) == 0 {
		return emptyInput
	}
	cmd := fields[0]

	if cmd == "exit" {
		return needExit
	}
	if cmd == "add" {
		return s.cmdAdd(fields)
	}
	if cmd == "list" {
		return s.cmdList(fields)
	}
	if cmd == "done" {
		return s.cmdDone(fields)
	}
	if cmd == "del" {
		return s.cmdDone(fields)
	}
	if cmd == "help" {
		return s.cmdHelp(fields)
	}
	if cmd == "events" {
		return s.cmdEvent(fields)
	}
	return unknownCommand

}
func (s *Scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return wrongArgs
	}
	title := fields[1]

	taskText := ""
	for i := 2; i < len(fields); i++ {
		taskText += fields[i]
		if i != len(fields)-1 {
			taskText += " "
		}
	}
	task := komandi.NewTask(title, taskText)

	s.todoList.Add(task)

	printAdd(title)

	return ""

}
func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}
	tasks := s.todoList.List()
	printTasks(tasks)
	return ""
}
func (s *Scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}
	title := fields[1]

	doneResult := s.todoList.Done(title)
	if doneResult != "" {
		return doneResult
	}
	printDone(title)
	return ""

}
func (s *Scanner) cmdDel(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}
	title := fields[1]

	delResult := s.todoList.Del(title)
	if delResult != "" {
		return delResult
	}
	printDel(title)
	return ""
}
func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}
	printHelp()
	return ""

}
func (s *Scanner) cmdEvent(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}
	printEvents(s.events)
	return ""

}
