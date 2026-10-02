package main

import (
	"fmt"
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

	for i, expense := range expenses {
		fmt.Printf("%d. %s - Rs%.2f\n", i+1, expense.Description, expense.Amount)
	}
	
}
