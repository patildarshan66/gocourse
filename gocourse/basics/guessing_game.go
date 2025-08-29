package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Guess Target
	target := random.Intn(100) + 1

	// Welcome User
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	var guess int

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println("Congratulations! You guessed the number!")
			break
		} else if guess < target {
			fmt.Println("Your guess is too low. Try again!")
		} else {
			fmt.Println("Your guess is too high. Try again!")
		}
	}

}
