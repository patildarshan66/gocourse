package advanced

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go func() {
		channel <- "Hello from channel"
		channel <- "Another message"
		for _, e := range "Golang" {
			channel <- "Alphabet: " + string(e)
		}
	}()

	// go func() {
	// 	message := <-channel
	// 	println(message)
	// 	message = <-channel
	// 	println(message)

	// }()

	message := <-channel
	println(message)
	message = <-channel
	println(message)

	for range 5 {
		msg := <-channel
		println(msg)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("End of main function")

}
