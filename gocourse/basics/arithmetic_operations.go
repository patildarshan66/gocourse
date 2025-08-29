package basics

import (
	"fmt"
	"math"
)

func main() {

	var a, b int = 10, 3
	var c, d float64 = 10, 3.0

	var result int
	var resultFloat float64

	// Addition
	result = a + b
	fmt.Println("Addition:", result)

	// Subtraction
	result = a - b
	fmt.Println("Subtraction:", result)

	// Multiplication
	result = a * b
	fmt.Println("Multiplication:", result)

	result = a % b
	fmt.Println("Modulus:", result)

	// Division
	resultFloat = c / d
	fmt.Println("Division:", resultFloat)

	// Overflow with signed integers
	var maxInt int64 = math.MaxInt64 //9223372036854775807
	fmt.Println("Max Int64:", maxInt)
	maxInt = maxInt + 1
	fmt.Println("Overflowed Int64:", maxInt)

	// Overflow with unsigned integers
	var maxUint uint64 = math.MaxUint64 // 18446744073709551615
	fmt.Println("Max Uint64:", maxUint)
	maxUint = maxUint + 1
	fmt.Println("Overflowed Uint64:", maxUint)

	// Underflow with signed integers
	var minFloat64 float64 = math.SmallestNonzeroFloat64 //5e-324
	fmt.Println("Min float64:", minFloat64)
	minFloat64 = minFloat64 / math.MaxFloat64
	fmt.Println("Underflowed float64:", minFloat64)

}
