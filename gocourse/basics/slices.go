package basics

import (
	"fmt"
	"slices"
)

func main() {

	var numbers []int
	fmt.Println(numbers)

	var numbers1 = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers1)

	numbers2 := []int{6, 7, 8, 9, 10}
	fmt.Println(numbers2)

	var slices3 []int = []int{1, 2, 4}
	fmt.Println(slices3)

	var slice1 []int = make([]int, 7)
	copy(slice1, []int{1, 2, 3, 4, 5, 6, 9})
	fmt.Println(slice1)

	slices3 = append(slices3, 5, 6, 7, 8, 9)
	fmt.Println(slices3)

	a := [5]int{1, 2, 3, 4, 5}
	newSlice := a[1:4]
	fmt.Println(newSlice)

	newSlice = append(newSlice, 10, 15, 20)
	fmt.Println(newSlice)

	var nilSlice []int
	fmt.Println(nilSlice)

	fmt.Println("Index 3 of nilSlice:", newSlice[3])
	newSlice[3] = 200
	fmt.Println("After changing index 3 of nilSlice:", newSlice)

	//check is two slices equal

	equalSlice := []int{1, 2, 3, 4, 5}
	equalSlice1 := []int{1, 2, 3, 4, 5}

	if slices.Equal(equalSlice, equalSlice1) {
		fmt.Println("Two slices are equal")
	}

	//Multi-dimensional slice
	var matrix = make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		matrix[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println(matrix)

	// creating slice from another slice using range
	mainSlice := slice1[2:4]
	fmt.Println("slice1:", slice1)
	fmt.Println("New Slice from slice1:", mainSlice)

	fmt.Println("Length of mainSlice:", len(mainSlice))
	fmt.Println("Capacity of mainSlice:", cap(mainSlice))

	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))

}
