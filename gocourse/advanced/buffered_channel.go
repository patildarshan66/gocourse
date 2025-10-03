package advaced

import (
	"fmt"
	"time"
)

func main() {
	/// ==============Example 1: BLOCKING ON SEND ONLY IF THE BUFFER IS FULL ================
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()
	fmt.Println("Blocking start here")
	ch <- 3 //blocking because the buffer is full
	fmt.Println("Blocking end here")
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)

	fmt.Println("End of program")

}

// func main() {

// 	///==============Example 2: BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY ================

// 	ch := make(chan int, 2)
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()

// 	fmt.Println("Received:", <-ch)
// 	fmt.Println("Received:", <-ch)
// 	fmt.Println("End of program")

// }
