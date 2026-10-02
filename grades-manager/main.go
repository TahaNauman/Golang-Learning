package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []float64
}

func calculateAverage(grades []float64) float64 {

	if len(grades) == 0 {
		return 0.0
	}

	total := 0.0
	for _, grade := range grades {
		total += grade
	}

	return total / float64(len(grades))
}

func findHighestGrade(grades []float64) (highest float64, index int) {
	if len(grades) == 0 {
		return 0.0, -1
	}

	highest = grades[0]
	index = 0

	for i, grade := range grades {
		if grade > highest {
			highest = grade
			index = i
		}
	}

	return highest, index
}

func findLowestGrade(grades []float64) (lowest float64, index int) {
	if len(grades) == 0 {
		return 0.0, -1
	}

	lowest = grades[0]
	index = 0

	for i, grade := range grades {
		if grade < lowest {
			lowest = grade
			index = i
		}
	}

	return lowest, index
}

func printStudentInfo(student Student) {
	fmt.Printf("Student: %s\n", student.Name)
	fmt.Println("Grades:")
	for _, grade := range student.Grades {
		fmt.Printf("  %.2f\n", grade)
	}
	average := calculateAverage(student.Grades)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Println()

	highest, index := findHighestGrade(student.Grades)
	if index != -1 {
		fmt.Printf("Highest Grade: %.2f (Index: %d)\n", highest, index)
	} else {
		fmt.Println("No grades available.")
	}

	lowest, index := findLowestGrade(student.Grades)
	if index != -1 {
		fmt.Printf("Lowest Grade: %.2f (Index: %d)\n", lowest, index)
	} else {
		fmt.Println("No grades available.")
	}
}

func main() {
	students := []Student{
		{Name: "Alice", Grades: []float64{85.5, 90.0, 78.0}},
		{Name: "Bob", Grades: []float64{92.0, 88.5, 95.0}},
		{Name: "Charlie", Grades: []float64{70.0, 75.5, 80.0}},
	}
	for _, student := range students {
		printStudentInfo(student)
	}

}
