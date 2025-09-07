package intermediate

import "fmt"

func main() {
	rect := Rectangle{length: 10, width: 9}
	area := rect.area()
	fmt.Println("Before Scale: ", area)
	rect.scale(2)
	area = rect.area()
	fmt.Println("After Scale:", rect.area())

	num := MyInt(10)
	num1 := MyInt(-10)

	fmt.Println("Is num is positive:", num.isPositive())
	fmt.Println("Is num1 is positive:", num1.isPositive())
	fmt.Println(num.welcomeMsg())

	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.area())
}

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

func (m MyInt) isPositive() bool {
	return m > 0
}

func (MyInt) welcomeMsg() string {
	return "Hello, there"
}
