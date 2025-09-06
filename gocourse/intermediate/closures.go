package intermediate

import "fmt"

func main() {

	incermenter := makeIncermentetr()
	fmt.Println(incermenter(1))
	fmt.Println(incermenter(2))
	fmt.Println(incermenter(3))
	fmt.Println(incermenter(4))
	fmt.Println(incermenter(5))
	fmt.Println(incermenter(6))
	fmt.Println(incermenter(7))
	fmt.Println(incermenter(8))
	fmt.Println(incermenter(9))
	fmt.Println(incermenter(10))

	incermenter2 := makeIncermentetr()
	fmt.Println(incermenter2(1))
	fmt.Println(incermenter2(2))
	fmt.Println(incermenter2(3))
	fmt.Println(incermenter2(4))
	fmt.Println(incermenter2(5))
	fmt.Println(incermenter2(6))
	fmt.Println(incermenter2(7))
	fmt.Println(incermenter2(8))
	fmt.Println(incermenter2(9))
	fmt.Println(incermenter2(10))

}

func makeIncermentetr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
