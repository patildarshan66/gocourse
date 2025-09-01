package basics

func main() {
	result := sum(1, 2, 3, 4, 5)
	println("Sum:", result)

	s, sum := add("Sum of 1,2,3,4 =", 1, 2, 3, 4)
	println(s, sum)

	numbers := []int{1, 2, 3, 4, 9}
	s1, sum1 := add("Sum of 1,2,3,4,9 =", numbers...)
	println(s1, sum1)
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func add(returnString string, numbers ...int) (string, int) {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return returnString, total
}
