package intermediate

import "fmt"

func main() {

	e := employee{
		person:     person{name: "Darshan", age: 26},
		employeeId: "284826",
		salary:     10000.00,
	}

	fmt.Println("Name:", e.name)
	fmt.Println("Age:", e.age)
	fmt.Println("EmployeeId:", e.employeeId)
	fmt.Println("Salary:", e.salary)

	e.introduce()

}

type person struct {
	name string
	age  int
}

type employee struct {
	person     //Embedded struct
	employeeId string
	salary     float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I am %s and I am %d year old\n", p.name, p.age)
}

func (e employee) introduce() {
	fmt.Printf("Hi, I am %s, Employee Id is %s and I earn %.2f\n", e.name, e.employeeId, e.salary)
}
