package basics

import (
	"errors"
	"fmt"
)

func main() {

	q, r := division(10, 3)

	fmt.Printf("Quotient: %d, Remainder: %v \n", q, r)

	result, err := compare(10, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison Result:", result)
	}

	name, age := getUserNameAge()
	fmt.Printf("User Name: %s and Age: %d\n", name, age)
}

func division(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func getUserNameAge() (name string, age int) {
	name = "Alice"
	age = 30
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare")
	}
}
