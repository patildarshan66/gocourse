package advanced

import (
	"fmt"
	"time"
)

func main() {

	var err error

	go sayHello()

	go printNumbers()

	go printLetters()
	go func() {
		err = doWork()
	}()
	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println("Error from goroutine:", err)
	} else {
		fmt.Println("No error from goroutine")
	}
	fmt.Println("Main function finished")
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, World!")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "Goroutines" {
		fmt.Println("Letter:", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occurred")
}
