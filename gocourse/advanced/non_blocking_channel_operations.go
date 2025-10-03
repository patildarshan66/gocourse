package advanced

import "time"

func main() {

	// ch := make(chan int)

	// ///===== NON BLOCKING RECEIVE OPERATION =====///

	// select {
	// case val := <-ch:
	// 	println("Received value:", val)
	// default:
	// 	println("No value received")
	// }

	///===== NON BLOCKING SEND OPERATION =====///

	// select {
	// case ch <- 42:
	// 	println("Sent value: 42")
	// default:
	// 	println("No value sent")
	// }

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case val := <-data:
				println("Received value:", val)
			case <-quit:
				println("Goroutine exiting")
				return
			default:
				println("No value received")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 * time.Second)
	}

	quit <- true
}
