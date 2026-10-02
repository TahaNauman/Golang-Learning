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

func findGradeRange(highest, lowest float64) float64 {
	return highest - lowest
}

func findGradeLetter(grade float64) (letter string) {
	switch {
	case grade >= 90:
		letter = "A"
	case grade >= 80:
		letter = "B"
	case grade >= 70:
		letter = "C"
	case grade >= 60:
		letter = "D"
	default:
		letter = "F"
	}

	return letter
}

func printStudentInfo(student Student) {
	fmt.Printf("Student: %s\n", student.Name)
	fmt.Println("Grades:")
	for _, grade := range student.Grades {
		letter := findGradeLetter(grade)
		fmt.Printf("  %.2f (%s)\n", grade, letter)
	}

	average := calculateAverage(student.Grades)
	overallLetter := findGradeLetter(average)
	fmt.Printf("Overall Average: %.2f (%s)\n", average, overallLetter)
	fmt.Println()

	highest, highestIndex := findHighestGrade(student.Grades)
	if highestIndex != -1 {
		fmt.Printf("Highest Grade: %.2f (Index: %d)\n", highest, highestIndex)
	} else {
		fmt.Println("No grades available.")
	}

	lowest, lowestIndex := findLowestGrade(student.Grades)
	if lowestIndex != -1 {
		fmt.Printf("Lowest Grade: %.2f (Index: %d)\n", lowest, lowestIndex)
	} else {
		fmt.Println("No grades available.")
	}

	if highestIndex != -1 && lowestIndex != -1 {
		fmt.Printf("Grade Range: %.2f\n", findGradeRange(highest, lowest))
	} else {
		fmt.Println("No grades available to calculate range.")
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
