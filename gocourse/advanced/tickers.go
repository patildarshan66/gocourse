package advanced

import "time"

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			println("Tick at", t.String())
		case <-stop:
			println("Stopping ticker")
			return
		}
	}
}

// ================== SCHEULING LOGGING, PERODIC TASKS, POLLING FOR UPDATES ==================
// func perodicTask() {
// 	fmt.Println("Task executed at:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			perodicTask()
// 		}
// 	}
// }

// func main() {

// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	i := 0
// 	for range 5 {
// 		i++
// 		fmt.Println(i)
// 	}

// 	for tick := range ticker.C {
// 		fmt.Println(tick)
// 	}

// }
