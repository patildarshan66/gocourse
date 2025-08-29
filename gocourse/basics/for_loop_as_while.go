package basics

import "fmt"

func main() {

	// for loop as a while
	sum := 0
	for {
		sum += 10
		fmt.Println("Sum:", sum)
		if sum == 100 {
			break
		}
	}
}
