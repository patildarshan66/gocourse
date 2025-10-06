package advanced

import (
	"context"
	"log"
	"time"
)

func doWork(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			// Context has been canceled or timed out
			println("Work canceled:", ctx.Err().Error())
			return
		default:
			// Simulate doing some work
			println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // timer of the context start here. You have this specific time duration to use this context.
	// After this time duration, the context will be canceled automatically.
	// defer cancel()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel() // Manually cancel the context after 2 seconds
	}()

	ctx = context.WithValue(ctx, "requestId", "request1234")
	go doWork(ctx)

	time.Sleep(3 * time.Second)
	requestId := ctx.Value("requestId")
	if requestId != nil {
		println("Request ID:", requestId.(string))
	} else {
		println("No Request ID found")
	}

	logWithContext(ctx, "This is a log message with context")
}

func logWithContext(ctx context.Context, message string) {
	requestId := ctx.Value("requestId")
	log.Printf("[RequestID: %v] %s", requestId, message)
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is Even", num)
// 		} else {
// 			return fmt.Sprintf("%d is Odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()
// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result for context.TODO():", result)

// 	// Or you can use context.Background
// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, time.Second)
// 	defer cancel()
// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result for context.Background():", result)

// 	time.Sleep(2 * time.Second)

// 	result = checkEvenOdd(ctx, 20)
// 	fmt.Println("Result after Timeout:", result)

// }
