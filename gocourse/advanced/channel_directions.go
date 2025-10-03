package advanced

func main() {

	ch := make(chan int)

	go producer(ch)

	consumer(ch)

}

func producer(ch chan<- int) { ///send only channel
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) { ///receive only channel
	for val := range ch {
		println("Received:", val)
	}
}
