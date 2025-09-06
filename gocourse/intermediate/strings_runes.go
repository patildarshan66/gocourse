package intermediate

import "fmt"

func main() {

	msg := "Hello,\nDarshan"
	rowMsg := `Hello\nWorld!`
	msg2 := "Hello, \tWorld!"
	msg3 := "Hello \rDarshan"
	msg4 := "Hello, \"Darshan\""
	msg5 := "Hello \n\rDarshan"

	fmt.Println(msg)
	fmt.Println(rowMsg)
	fmt.Println(msg2)
	fmt.Println(msg3)
	fmt.Println(msg4)
	fmt.Println(msg5)

	fmt.Println("Length of msg is:", len(msg))       //\n is considered as 1 character
	fmt.Println("Length of rowMsg is:", len(rowMsg)) //\n is considered as 2 characters

	fmt.Println("The first character ASCII is:", msg[0])          //ASCII value of H
	fmt.Println("The first character String is:", string(msg[0])) //String value of H

	///concatenation of strings

	greeting := "Hello"
	name := "Darshan"
	completeMsg := greeting + ", " + name + "!"
	fmt.Println(completeMsg)

	// compare strings
	str1 := "apple"
	str2 := "Apple"
	str3 := "Banana"
	str4 := "banana"
	str5 := "app"

	fmt.Println(str1 > str2)
	fmt.Println(str2 > str3)
	fmt.Println(str3 > str4)
	fmt.Println(str4 > str5)
	fmt.Println(str1 > str5)

	for _, char := range msg {
		fmt.Printf("%c\n", char) //%c is used to print character
		fmt.Printf("%U\n", char) //%U is used to print Unicode
		fmt.Printf("%d\n", char) //%d is used to print ASCII value
		fmt.Printf("%x\n", char) //%x is used to print hexadecimal value
		fmt.Println("-----")
	}

	var runeStr rune = 'a'      //rune is used to store a single character
	fmt.Println(runeStr)        //prints the ASCII value of 'a'
	fmt.Printf("%c\n", runeStr) //prints the character 'a'

	cstr := string(runeStr) //converting rune to string
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is%T\n", cstr) //prints the type of cstr

}
