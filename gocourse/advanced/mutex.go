package advanced

import "sync"

func main() {
	var wg sync.WaitGroup

	var mutex sync.Mutex
	counter := 0
	numGoroutines := 10

	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	println("Final count:", counter)

}

// type counter struct {
// 	mutex sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()
// 	return c.count
// }

// func main() {
// 	var wg sync.WaitGroup
// 	c := &counter{}
// 	numGoroutines := 10

// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				c.increment()
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final count:", c.getValue())
// }
