package intermediate

import (
	"fmt"
	"regexp"
)

func main() {

	//compile regular rexpression pattern to match valid email address
	reg := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	email1 := "darshanpatil66@gmail.com"
	email2 := "invalidemail@c"

	//Match Check
	fmt.Println(reg.MatchString(email1))
	fmt.Println(reg.MatchString(email2))

	//Capturing groups
	// compile a regex  pattern to capture  date components
	re := regexp.MustCompile(`(\d{2})/(\d{2})/(\d{4})`)
	re1 := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "26/02/1999"
	date2 := "1999-02-26"

	subGroupes := re.FindStringSubmatch(date)
	subGroupes1 := re1.FindStringSubmatch(date2)

	fmt.Println(subGroupes)
	fmt.Println(subGroupes[0])
	fmt.Println(subGroupes[1])
	fmt.Println(subGroupes[2])
	fmt.Println(subGroupes[3])
	fmt.Println("--------------------")
	fmt.Println(subGroupes1)
	fmt.Println(subGroupes1[0])
	fmt.Println(subGroupes1[1])
	fmt.Println(subGroupes1[2])
	fmt.Println(subGroupes1[3])

	// Traget string replacing chars
	str := "Hello World"
	reg3 := regexp.MustCompile(`[aeiou]`)

	result := reg3.ReplaceAllString(str, "#")

	fmt.Println(result)

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`) // match lower case occurencess
	testStr := "Golang is great going"
	fmt.Println("Match:", re.MatchString(testStr))

}
