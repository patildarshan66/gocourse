package intermediate

import "fmt"

func main() {

	num := 234
	fmt.Printf("%05d\n", num)
	num1 := 2347879898
	fmt.Printf("%05d\n", num1)

	msg := "Hello"

	fmt.Printf("|%10s|\n", msg)
	fmt.Printf("|%-10s|\n", msg)

	hello1 := "Hello \nWord!"
	hello2 := `Hello \nWord!`

	fmt.Println(hello1)
	fmt.Println(hello2)

}
