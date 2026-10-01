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

		var choice int
		fmt.Println("")
		fmt.Println("=== To Do List ===")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())

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
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}
