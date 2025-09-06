package intermediate

import "fmt"

func main() {

	var x int = 20

	//address of operator x
	fmt.Println(&x)

	//1st way of declaring pointer
	var b = &x //refrence operator
	fmt.Println(b)
	fmt.Println(*b) //dereference operator
	//address of operator b
	fmt.Println(&b)

	//2nd way of declaring pointer
	var c *int = &x
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(&c)

	///nil pointer
	var d *int

	fmt.Println(d)

	// incrementing value using pointer
	makeIncremeter(b)
	fmt.Println(x)
	makeIncremeter(c)
	fmt.Println(x)
}

func makeIncremeter(ptr *int) {
	*ptr++
}
