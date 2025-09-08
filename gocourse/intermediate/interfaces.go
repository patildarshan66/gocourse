package intermediate

import (
	"fmt"
	"math"
)

func main() {

	c := circle{radius: 10}
	r := rect{height: 10, width: 10}
	// r2 := rect2{height: 10,width: 10}

	measure(c)
	fmt.Println(c.diameter())
	measure(r)
	// measure(r2) // Giving error beacuse of geometery interface all methods are not implemented by rect2

	printType(9)
	printType("Darshan")
	printType(false)

	myPrinter(1, 2, "Darshan", false, 1.2)

}

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

type rect2 struct {
	width, height float64
}

func (r rect2) area() float64 {
	return r.height * r.width
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func printType(a interface{}) {
	switch a.(type) {

	case int:
		fmt.Println("A is type of integer")
	case string:
		fmt.Println("A is type of string")
	default:
		fmt.Println("A is type of Unknown")
	}
}

func myPrinter(a ...interface{}) {
	for _, i := range a {
		fmt.Println(i)
	}
}
