package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	str := "12345"
	a, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error Parsing value:", err)
		return
	}

	fmt.Println("Parsed Number is :", a)

	b, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		fmt.Println("Error Parsing value:", err)
		return
	}
	fmt.Println("Parsed Number is :", b)

	str2 := "3.14"

	c, err := strconv.ParseFloat(str2, 64)

	if err != nil {
		fmt.Println("Error Parsing value:", err)
		return
	}
	fmt.Printf("Parsed Number is : %.2f\n", c)

	binaryStr := "1010" // 0 + 2 + 0 + 8 =10

	d, err := strconv.ParseInt(binaryStr, 2, 64)

	if err != nil {
		fmt.Println("Error Parsing value:", err)
		return
	}
	fmt.Println("Parsed Binary to Decimal is :", d)

	hexStr := "FF"
	e, err := strconv.ParseInt(hexStr, 16, 64)

	if err != nil {
		fmt.Println("Error Parsing value:", err)
		return
	}
	fmt.Println("Parsed Hex to Decimal Number is :", e)

	invalidNum := "123bdkd"

	f, err := strconv.ParseInt(invalidNum, 10, 32)

	if err != nil {
		fmt.Println("Error Parsing value:", err)
		return
	}
	fmt.Println("Parsed Number is :", f)

}
