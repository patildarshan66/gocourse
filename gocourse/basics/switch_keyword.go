package basics

import "fmt"

func main() {

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Sunday":
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("Invalid day.")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or greater")
	}

	num := 2

	switch {
	case num > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is equal to 2")
	default:
		fmt.Println("Not two")
	}

	checkType(9)
	checkType("Hello")
	checkType(true)
	checkType(3.14)

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Unknown type")
	}
}
