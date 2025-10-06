package advanced

import (
	"fmt"
	"time"
)

/// ====== Multiple Timers ======

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	timer2 := time.NewTimer(4 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Timer 1 expired after 2 seconds")
	case <-timer2.C:
		fmt.Println("Timer 2 expired after 4 seconds")
	}
}

///========Scheduling Delayed Operations ========

// func main() {

// 	timer := time.NewTimer(2 * time.Second)
// 	go func() {
// 		<-timer.C
// 		fmt.Println("Timer expired after 2 seconds")
// 	}()
// 	fmt.Println("waiting....")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Done")
// }

///============Timeout example============

// func longRunningTask() {
// 	for i := range 20 {
// 		fmt.Println("Task running:", i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {

// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)

// 	go func() {
// 		longRunningTask()
// 		done <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Timeout! Task took too long.")
// 	case <-done:
// 		fmt.Println("Task completed successfully.")
// 	}
// }

///=============Basic timer example============

// func main() {
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}
// 	timer.Reset(1 * time.Second)
// 	fmt.Println("Timer reset")
// 	<-timer.C
// 	fmt.Println("Timer expired")
// }
