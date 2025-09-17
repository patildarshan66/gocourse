package intermediate

import (
	"flag"
	"fmt"
)

func main() {

	// fmt.Println("Command:", os.Args[0])

	// for i, args := range os.Args {
	// 	fmt.Println("Command[", i, "]:", args)
	// }

	//Define flags

	var name string
	var age int
	var isMale bool

	flag.StringVar(&name, "name", "Darshan", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&isMale, "male", false, "Gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("IsMale:", isMale)

}
