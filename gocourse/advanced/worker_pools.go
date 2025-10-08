package advanced

import (
	"fmt"
	"time"
)

// / SAMULATING TICKET BOOKING SYSTEM USING WORKER POOLS

type ticketRequest struct {
	personId   int
	numTickets int
	cost       int
}

func ticketProccessor(ticketRequests <-chan ticketRequest, results chan<- int) {
	for request := range ticketRequests {
		fmt.Printf("Processing ticket booking for person %d: %d tickets with the cost of %d\n", request.personId, request.numTickets, request.cost)
		time.Sleep(time.Second) // simulate time taken to process the ticket booking
		results <- request.personId
	}
}

func main() {

	numWorkers := 3
	numRequests := 5
	ticketCost := 100

	ticketRequests := make(chan ticketRequest, numRequests)
	results := make(chan int, numRequests)

	// Create worker goroutines
	for range numWorkers {
		go ticketProccessor(ticketRequests, results)
	}

	for i := range numRequests {
		numTickets := (i + 1) * 2
		ticketRequests <- ticketRequest{
			personId:   i + 1,
			numTickets: numTickets,
			cost:       numTickets * ticketCost,
		}
	}

	close(ticketRequests)

	for range numRequests {
		personId := <-results
		fmt.Printf("Ticket booking completed for person %d\n", personId)
	}

}

///============================== Basic Worker Pools Pattern ==========================

// func worker(id int, tasks <-chan int, results chan<- string) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		// Simulate doing some work
// 		time.Sleep(time.Second)
// 		results <- fmt.Sprintf("Worker %d finished task %d", id, task)
// 	}
// }

// func main() {

// 	numWorkers := 3
// 	numJobs := 10

// 	/// 10 task will be distributed among 3 workers based of first complete first serve basis

// 	tasks := make(chan int, numJobs)
// 	results := make(chan string, numJobs)

// 	//create workers
// 	for w := 1; w <= numWorkers; w++ {
// 		go worker(w, tasks, results)
// 	}

// 	//send tasks to workers
// 	for i := 1; i <= numJobs; i++ {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	//collect results from workers
// 	for i := 1; i <= numJobs; i++ {
// 		result := <-results
// 		fmt.Println(result)
// 	}
// }
