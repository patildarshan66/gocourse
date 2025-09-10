package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	strN := "Hello"
	strM := "World"
	result := strN + " " + strM
	fmt.Println(result)
	fmt.Println(result[1:8])
	fmt.Println(strM[0])

	// Int to String conversion
	num := 18
	str := strconv.Itoa(num)
	fmt.Println(len(str))

	//String spliting
	str1 := "Darshan Patil"
	splitResult := strings.Split(str1, " ")
	fmt.Println(splitResult)

	// String joining
	fmt.Println(strings.Join(splitResult, "_"))

	//String contains
	fmt.Println(strings.Contains(str1, "Darshan"))
	fmt.Println(strings.Contains(str1, "darshan"))

	//String Replace
	repString := "My name is Darshan Patil, My self Darshan Patil, I am Darshan"
	fmt.Println(strings.Replace(repString, "Darshan", "Kalpesh", 1))  //Replace first one occurence
	fmt.Println(strings.Replace(repString, "Darshan", "Kalpesh", 2))  //Replace first two occurences
	fmt.Println(strings.Replace(repString, "Darshan", "Kalpesh", -1)) //Replace all occurences
	fmt.Println(strings.ReplaceAll(repString, "Darshan", "Kalpesh"))  //Replace all occurences

	//String Trim
	trimStr := " Darshan Patil "
	fmt.Println(trimStr)
	fmt.Println(strings.TrimSpace(trimStr))

	//String toLower, toUpper
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToUpper(str1))

	// String Repeat
	fmt.Println(strings.Repeat("Darshan ", 5))

	//String count
	fmt.Println(strings.Count(str1, "a"))

	//String Prefix and Sufix check
	fmt.Println(strings.HasPrefix(str1, "Dar"))
	fmt.Println(strings.HasPrefix(str1, "dar"))
	fmt.Println(strings.HasSuffix(str1, "til"))
	fmt.Println(strings.HasSuffix(str1, "Til"))

	// Regular Expression
	rexString := "Darshan 1234 Patil 8908, Pincode 700706"
	regx := regexp.MustCompile(`\d+`)
	fmt.Println(regx.FindAllString(rexString, -1)) // find all matches -1
	fmt.Println(regx.FindAllString(rexString, 2))  //find 2 matches

	// rune count
	runeChar := "Hello, こんにちは"
	fmt.Println(utf8.RuneCountInString(runeChar))
	fmt.Println(len(runeChar))

	//define String builder
	var builder strings.Builder

	//write string into builder buffer
	builder.WriteString("Darshan")
	builder.WriteString("Patil")
	//write rune in string buffer
	builder.WriteRune(' ')
	builder.WriteString("At Post: ")
	builder.WriteString("Mhasawad")

	//convert builder to string
	strResult := builder.String()
	fmt.Println(strResult)
	//write string into exsiting builder buffer
	builder.WriteString(" Pincode: 425432")

	//convert builder to string after
	strResult = builder.String()
	fmt.Println(strResult)
	// Reset the builder buffer //clear the builder
	builder.Reset()

	//write string into builder buffer
	builder.WriteString("Kalpesh Patil")
	//convert builder to string after
	strResult = builder.String()
	fmt.Println(strResult)

}
