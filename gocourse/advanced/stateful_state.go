package advanced

import (
	"fmt"
	"time"
)

type StatefulState struct {
	counter int
	ch      chan int
}

func (s *StatefulState) start() {
	go func() {
		for {
			select {
			case value := <-s.ch:
				s.counter += value

				fmt.Println("Current Count:", s.counter)
			}
		}
	}()
}

func (s *StatefulState) send(value int) {
	s.ch <- value
}

func main() {
	stWorker := &StatefulState{ch: make(chan int)}

	stWorker.start()

	for i := 0; i < 5; i++ {
		stWorker.send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
