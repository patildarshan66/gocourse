package advanced

import (
	"fmt"
	"time"
)

// // /=================== Synchronization Multiple Goroutines and Ensuring that all Goroutines are complete ===================///
// func main() {

// 	numGoroutines := 3
// 	done := make(chan bool, numGoroutines)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working....\n", id)
// 			time.Sleep(2 * time.Second)
// 			done <- true
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done
// 	}

// 	fmt.Println("All goroutines complete.")
// }

// / Synchronization data exchange
func main() {

	ch := make(chan string)

	go func() {
		for i := range 5 {
			ch <- "Hello " + string(rune('0'+i))
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()
	for value := range ch {
		fmt.Println("Received:", value, ":", time.Now())
	}
}
