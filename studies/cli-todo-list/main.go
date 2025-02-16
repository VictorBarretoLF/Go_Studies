package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text string
	Done bool
}

var tasks []Task

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose an option: add, list, complete, remove, exit\n")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		switch input {
		case "add":
			addTaskHandler(scanner)
		case "list":
			listTasksHandler()
		case "complete":
			completeTaskHandler(scanner)
		case "remove":
			removeTaskHandler(scanner)
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, try again!")
		}
	}
}

func addTaskHandler(scanner *bufio.Scanner) {
	fmt.Print("Enter task: ")
	scanner.Scan()
	taskText := scanner.Text()
	tasks = append(tasks, Task{Text: taskText, Done: false})
	fmt.Println("Task added!")
}

func listTasksHandler() {
	fmt.Println("\nTo-Do List:")

	if hasNoTasks(tasks) {
		fmt.Println("No tasks yet!")
		return
	}

	for i, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}
}

func completeTaskHandler(scanner *bufio.Scanner) {
	fmt.Print("Enter task number to mark as done: ")

	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil || isInvalidIndex(num, tasks) {
		fmt.Println("Invalid number!")
		return
	}

	tasks[num-1].Done = true
	fmt.Println("Task marked as completed!")
}

func removeTaskHandler(scanner *bufio.Scanner) {
	fmt.Print("Enter task number to remove: ")

	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil || isInvalidIndex(num, tasks) {
		fmt.Println("Invalid number!")
		return
	}

	tasks = append(tasks[:num-1], tasks[num:]...)
	fmt.Println("Task removed!")
}

func isInvalidIndex(index int, tasks []Task) bool {
	return index < 1 || index > len(tasks)
}

func hasNoTasks(tasks []Task) bool {
	return len(tasks) == 0
}