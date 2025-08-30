package basics

import "fmt"

func main() {
	result := add(5, 10)
	fmt.Println("Result:", result)

	greet := func() {
		fmt.Println("Hello, World!")
	}
	greet()

	operation := add

	fmt.Println("Operation Result:", operation(3, 7))

	resultOp := applyOperation(10, 10, add)
	fmt.Println("Apply Operation Result:", resultOp)

	mutipler := makeMultiplier(5)
	fmt.Println("Multipler:", mutipler(5))
}

func add(a int, b int) int {
	return a + b
}

func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
