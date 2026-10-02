package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Expense struct {
	Amount      float64
	Description string
}

func main() {

	expenses := []Expense{
		{Amount: 50.0, Description: "Groceries"},
		{Amount: 20.0, Description: "Transport"},
		{Amount: 100.0, Description: "Utilities"},
	}
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("===Expense Tracker===")
		fmt.Println("1. View Expenses")
		fmt.Println("2. Add Expense")
		fmt.Println("3. Show Total Expenses")
		fmt.Println("4. Delete Expense")
		fmt.Println("5. Edit Expense")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid option.")
			return
		}

		switch choice {
		case 1:
			fmt.Println("=== Expenses ===")
			for i, expense := range expenses {
				fmt.Printf("%d. %s - Rs%.2f\n", i+1, expense.Description, expense.Amount)
			}
		case 2:
			var amount float64
			var description string
			fmt.Println("=== Add Expense ===")
			fmt.Print("Enter expense amount: ")
			scanner.Scan()
			amount, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid amount.")
				continue
			}
			fmt.Print("Enter expense description: ")
			scanner.Scan()
			description = scanner.Text()
			expenses = append(expenses, Expense{Amount: amount, Description: description})
			fmt.Println("Expense added successfully.")
		case 3:
			total := 0.0
			for _, expense := range expenses {
				total += expense.Amount
			}
			fmt.Printf("Total Expenses: Rs%.2f\n", total)
			
		case 4:
			fmt.Println("=== Delete Expense ===")
			fmt.Print("Enter the index of the expense to delete: ")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid index.")
				continue
			}
			if index < 1 || index > len(expenses) {
				fmt.Println("Invalid index. Please enter a valid index.")
				continue
			}
			expenses = append(expenses[:index-1], expenses[index:]...)
			fmt.Println("Expense deleted successfully.")
		case 5:
			fmt.Println("=== Edit Expense ===")
			fmt.Print("Enter the index of the expense to edit: ")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid index.")
				continue
			}
			if index < 1 || index > len(expenses) {
				fmt.Println("Invalid index. Please enter a valid index.")
				continue
			}
			fmt.Println("What would you like to edit")
			fmt.Println("1. Expense Amount")
			fmt.Println("2. Expense Description")
			fmt.Print("Enter your choice: ")
			scanner.Scan()
			choice, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid choice.")
				continue
			}
			switch choice {
			case 1:
				fmt.Print("Enter new expense amount: ")
				scanner.Scan()
				newAmount, err := strconv.ParseFloat(scanner.Text(), 64)
				if err != nil {
					fmt.Println("Invalid input. Please enter a valid amount.")
					continue
				}
				expenses[index-1].Amount = newAmount
			case 2:
				fmt.Print("Enter new expense description: ")
				scanner.Scan()
				newDescription := scanner.Text()
				expenses[index-1].Description = newDescription
			default:
				fmt.Println("Invalid choice. Please try again.")
				continue
			}
			fmt.Println("Expense edited successfully.")
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}	

}
