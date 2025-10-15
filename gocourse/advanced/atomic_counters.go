package advanced

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounters struct {
	counter int64
}

func (ac *AtomicCounters) increment() {
	atomic.AddInt64(&ac.counter, 1)
}

func (ac *AtomicCounters) getValue() int64 {
	return atomic.LoadInt64(&ac.counter)
}

func main() {

	var wg sync.WaitGroup
	numGoRutines := 10
	counter := &AtomicCounters{}

	for range numGoRutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter Value:%d\n", counter.getValue())

}
