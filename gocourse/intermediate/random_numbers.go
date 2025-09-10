package intermediate

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// It will generate different bumbers on every time we run program.
	// Because of dynamic seed value assigned by go internally
	fmt.Println(rand.Intn(100) + 1)
	fmt.Println(rand.Intn(100) + 1)
	fmt.Println(rand.Intn(100) + 1)
	fmt.Println(rand.Intn(100) + 1)

	// It will generate same bumbers on every time we run program like.
	// Because of constant seed value
	// first run: 13 65 64 72 7 57
	// second run: 13 65 64 72 7 57
	val := rand.New(rand.NewSource(55))
	fmt.Println(val.Intn(100) + 1)
	fmt.Println(val.Intn(100) + 1)
	fmt.Println(val.Intn(100) + 1)
	fmt.Println(val.Intn(100) + 1)
	fmt.Println(val.Intn(100) + 1)
	fmt.Println(val.Intn(100) + 1)

	// It will generate different bumbers on every time we run program.
	// Because of we are giving seed value dynamic from time epoch
	val2 := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(val2.Intn(100) + 1)
	fmt.Println(val2.Intn(100) + 1)
	fmt.Println(val2.Intn(100) + 1)
	fmt.Println(val2.Intn(100) + 1)
	fmt.Println(val2.Intn(100) + 1)
	fmt.Println(val2.Intn(100) + 1)

	// random flot number
	fmt.Println(rand.Float64())

	for {
		fmt.Println("Welcome to the dice game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Println("Enter your choice (1 or 2):")

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, Please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		die1 := rand.Intn(6) + 1

		die2 := rand.Intn(6) + 1

		fmt.Printf("You rolled a %d and a %d\n", die1, die2)
		fmt.Println("Total:", die1+die2)

		// Ask of the user wants to roll again

		fmt.Println("Do you want to roll again? (y/n):")
		var rollAgain string
		_, err1 := fmt.Scan(&rollAgain)
		rollAgain = strings.ToLower(rollAgain)
		if err1 != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}

		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

	}

}
