package basics

import "fmt"

func main() {

	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println("Integer Array:", numbers)

	var namesArr = [3]string{"Darshan", "Kalpesh", "Rohit"}
	fmt.Println("String Array:", namesArr)

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray // This creates a copy of the array
	copiedArray[0] = 100         // Modifying the copied array
	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	originalArray2 := [3]int{4, 5, 6}
	copiedArray2 := &originalArray2 // Creating a pointer to the original array (passing address of originalArray2)
	copiedArray2[0] = 200           // Modifying the copied array which affects the original array
	fmt.Println("Original Array 2:", originalArray2)
	fmt.Println("Copied Array 2:", *copiedArray2)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Number at index", i, "is", numbers[i])
	}

	for i, v := range namesArr {
		fmt.Println("Name at index", i, "is", v)
	}

	age, name := someFunction()
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)

	// Ignoring first return value using underscore, Used to store unused return value
	_, userName := someFunction()
	fmt.Println("User Name:", userName)

	//comparing two arrays
	var array1 [3]int = [3]int{1, 2, 3}
	var array2 [3]int = [3]int{1, 2, 3}
	fmt.Println("Array 1 is equal to Array 2:", array1 == array2)

	//multidimensional array
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)
}

func someFunction() (int, string) {
	return 26, "Darshan"
}
