package intermediate

import "fmt"

func main() {

	// printing actions

	// fmt.Print("Hello ")
	// fmt.Print("World!")
	// fmt.Print(123, 455)

	// fmt.Println("Hello ")
	// fmt.Println("World!")
	// fmt.Println(123, 455)

	// name := "John"
	// age := 30

	// fmt.Printf("My name is %s and I am %d years old.\n", name, age)
	// fmt.Printf("Binary: %b, Hex: %x\n", age, age)

	// Formatting Functions
	// s := fmt.Sprint("Hello", "World!", 123, 455)
	// fmt.Print(s)

	// s1 := fmt.Sprintln("Hello", "World!", 123, 455)
	// fmt.Print(s1)
	// fmt.Print("new add here by Sprintln here.")

	// s2 := fmt.Sprintf("My name is %s and I am %d years old.\n", "John", 30)
	// fmt.Println(s2)
	// fmt.Println(s2)

	// Scanning Functions

	// var name1 string
	// var age1 int

	// fmt.Println("Enter your name and age: ")
	// fmt.Scan(&name1, &age1) // waits for the user to enter both name and age
	// fmt.Printf("Hello %s, you are %d years old.\n", name1, age1)

	// var name2 string
	// var age2 int
	// fmt.Println("Enter your name and age: ")
	// fmt.Scanln(&name2, &age2) // waits for the user to press enter
	// fmt.Printf("Hello %s, you are %d years old.\n", name2, age2)

	// var name3 string
	// var age3 int
	// fmt.Println("Enter your name and age: ")
	// fmt.Scanf("%s %d", &name3, &age3)  // waits for the user to enter both name and age in the specified format (Darshan 24)
	// // fmt.Scanf("%s, %d", &name3, &age3) // waits for the user to enter both name and age in the specified format (Darshan, 24)
	// fmt.Printf("Hello %s, you are %d years old.\n", name3, age3)

	// Error Formatting functions

	err := checkAge(17)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid.")
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is less than 18", age)
	}
	return nil
}
