package advanced

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		println("Received:", val)
	}

}

// / Ranged over closed Channel Example
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		println("Received:", val)
// 	}
// }

// / Receiving from a Closed Channel Example
// func main() {

// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch

// 	if !ok {
// 		println("Channel is closed")
// 	} else {
// 		println("Received:", val)
// 	}

// }

// / Simple Closing Channels Example
// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for range 5 {
// 		println(<-ch)
// 	}
// }
