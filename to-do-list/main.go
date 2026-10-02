package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Title     string
	Completed bool
}

func main() {
	tasks := []Task{
		{Title: "Buy groceries", Completed: false},
		{Title: "Clean the house", Completed: true},
		{Title: "Finish project", Completed: false},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("")
		fmt.Println("=== To Do List ===")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Edit Task")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("=== Tasks ===")
			fmt.Println("")
			for i, task := range tasks {
				fmt.Printf("%d. %s %s\n", i+1, map[bool]string{true: "[✓]", false: "[]"}[task.Completed], task.Title)
			}
		case 2:
			var title string
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title = scanner.Text()
			tasks = append(tasks, Task{Title: title, Completed: false})
		case 3:
			var taskNumber int
			fmt.Print("Enter task number to complete: ")
			scanner.Scan()
			taskNumber, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if taskNumber > 0 && taskNumber <= len(tasks) {
				tasks[taskNumber-1].Completed = true
				fmt.Println("Task marked as completed.")
			} else {
				fmt.Println("Invalid task number.")
			}
		case 4:
			var taskNumber int
			fmt.Print("Enter task number to delete: ")
			scanner.Scan()
			taskNumber, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if taskNumber > 0 && taskNumber <= len(tasks) {
				tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
				fmt.Println("Task deleted.")
			} else {
				fmt.Println("Invalid task number.")
			}
		case 5:
			var taskNumber int
			fmt.Print("Enter task number to edit: ")
			scanner.Scan()
			taskNumber, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if taskNumber > 0 && taskNumber <= len(tasks) {
				var newTitle string
				fmt.Print("Enter new task title: ")
				scanner.Scan()
				newTitle = scanner.Text()
				tasks[taskNumber-1].Title = newTitle
				fmt.Println("Task updated.")
			} else {
				fmt.Println("Invalid task number.")
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}
