package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Choose Difficulty: ")
	var difficulty string
	var secretNumber, upperLimit, maxAttempts int

	fmt.Println("1. Easy 1-50")
	fmt.Println("2. Medium 1-100")
	fmt.Println("3. Hard 1-500")
	fmt.Print("Choice: ")
	fmt.Scan(&difficulty)

	attempts := 0
	switch difficulty {
	case "1":
		upperLimit = 50
		maxAttempts = 10
	case "2":
		upperLimit = 100
		maxAttempts = 7
	case "3":
		upperLimit = 500
		maxAttempts = 5
	default:
		fmt.Println("Invalid choice. Defaulting to Medium difficulty.")
		upperLimit = 100
		maxAttempts = 7
	}

	secretNumber = rand.Intn(upperLimit) + 1

	for {

		var guess int
		var playAgain string

		if attempts >= maxAttempts {
			fmt.Println("You've used up all of your attempts")
			fmt.Println("Game over!")
			fmt.Printf("The secret number was: %d\n", secretNumber)
			fmt.Print("Do you want to play again? (y/n): ")
			fmt.Scan(&playAgain)

			if playAgain == "y" || playAgain == "Y" {
				attempts = 0
				secretNumber = rand.Intn(upperLimit) + 1
				continue
			} else {
				fmt.Println("Thanks for playing!")
				break
			}
		}

		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		if guess < 1 || guess > upperLimit {
			fmt.Printf("Please enter a number between 1 and %d.\n", upperLimit)
			continue
		}

		attempts++

		if guess > secretNumber {
			fmt.Println("Your guess is too high.")
		} else if guess < secretNumber {
			fmt.Println("Your guess is too low.")
		} else {
			fmt.Println("Congratulations! You guessed the number.")
			fmt.Printf("It took you %d attempts.\n", attempts)
			fmt.Print("Do you want to play again? (y/n): ")
			fmt.Scan(&playAgain)

			if playAgain == "y" || playAgain == "Y" {
				attempts = 0
				secretNumber = rand.Intn(upperLimit) + 1
				continue
			} else {
				fmt.Println("Thanks for playing!")
				break
			}
		}
	}
}
