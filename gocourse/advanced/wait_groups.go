package advanced

import (
	"fmt"
	"sync"
	"time"
)

/// CONSTRUCTION Real World Example ===================

type Worker struct {
	id   int
	task string
}

func (w *Worker) DoWork(wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate work
	fmt.Printf("Worker %d is working on %s\n", w.id, w.task)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d has completed %s\n", w.id, w.task)
}

func main() {
	var wg sync.WaitGroup

	tasks := []string{"Task A", "Task B", "Task C"}
	// workers := []Worker{
	// 	{id: 1, task: "Task A"},
	// 	{id: 2, task: "Task B"},
	// 	{id: 3, task: "Task C"},
	// }

	// wg.Add(len(workers))

	for i, task := range tasks {
		worker := Worker{id: i + 1, task: task}
		wg.Add(1)
		go worker.DoWork(&wg)
	}

	wg.Wait()
	fmt.Println("All workers have completed their tasks")
}

// /// Example USING CHANNELS With Result and Task Channel ===================

// func worker(id int, tasks <-chan int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// Simulate work
// 	fmt.Printf("Worker %d is working\n", id)
// 	time.Sleep(time.Second)
// 	for task := range tasks {
// 		result <- task * 2 // Send result to channel
// 	}

// 	fmt.Printf("Worker %d is done\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5

// 	result := make(chan int, numJobs) // Buffered channel to hold results
// 	tasks := make(chan int, numJobs)  // Buffered channel to hold tasks

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, result, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1 // Send tasks to channel
// 	}
// 	close(tasks) // Close tasks channel after sending all tasks
// 	go func() {
// 		wg.Wait()
// 		close(result) // Close the channel after all workers are done
// 	}()

// 	// Collect results
// 	for res := range result {
// 		fmt.Printf("Received result: %d\n", res)
// 	}

// 	fmt.Println("All workers completed")
// }

// /// Example USING CHANNELS With Result Channel ===================

// func worker(id int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// Simulate work
// 	fmt.Printf("Worker %d is working\n", id)
// 	time.Sleep(time.Second)
// 	result <- id * 2 // Send result to channel
// 	fmt.Printf("Worker %d is done\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 3

// 	result := make(chan int, numJobs) // Buffered channel to hold results

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, result, &wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(result) // Close the channel after all workers are done
// 	}()

// 	// Collect results
// 	for res := range result {
// 		fmt.Printf("Received result: %d\n", res)
// 	}

// 	fmt.Println("All workers completed")
// }

/// ============== BASIC EXAMPLE WITHOUT USING CHANNELS ==============

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// Simulate work
// 	fmt.Printf("Worker %d is working\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d is done\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers completed")
// }
