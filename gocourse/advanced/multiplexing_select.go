package advanced

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)

	go func() {
		ch1 <- 42
		close(ch1)
	}()

	for {
		select {
		case val, ok := <-ch1:
			if !ok {
				fmt.Println("Channel closed, exiting loop")
				return
			}
			fmt.Println("Received:", val)
		}
	}
}

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 42
// 	}()

// 	select {
// 	case val := <-ch:
// 		println("Received:", val)
// 	case <-time.After(1 * time.Second):
// 		println("Timeout: No data received within 1 second")
// 	}
// 	fmt.Println("End of Program")
// }

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- 42

// 	}()
// 	go func() {
// 		ch2 <- 84
// 	}()

// 	time.Sleep(2 * time.Second)
// 	for range 2 {
// 		select {
// 		case val1 := <-ch1:
// 			println("Received from ch1:", val1)
// 		case val2 := <-ch2:
// 			println("Received from ch2:", val2)
// 		default:
// 			println("No data received")
// 		}
// 	}

// 	fmt.Println("End of Program")
// }
