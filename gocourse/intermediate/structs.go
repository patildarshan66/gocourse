package intermediate

import "fmt"

func main() {
	p := Person{firstName: "John", lastName: "Doe", age: 30, address: Address{city: "Nandurbar", state: "Maharashtra"}, ContactInfo: ContactInfo{email: "john.doe@example.com", mobile: "123-456-7890"}}
	fmt.Println(p.fullName())
	fmt.Println("Before increasing age:", p.age)
	p.increaseAgeByOne()
	fmt.Println("After increasing age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
	fmt.Println("Email:", p.email)
	fmt.Println("Mobile:", p.mobile)

	p1 := Person{firstName: "Alice", lastName: "Johnson"}
	p2 := Person{firstName: "Alice", lastName: "Johnson"}
	p4 := Person{firstName: "Darshan", lastName: "Patil"}

	fmt.Println("P == P1:", p == p1)
	fmt.Println("P1 == P2:", p1 == p2)
	fmt.Println("P1 == P4:", p1 == p4)
	// Anonymous struct
	p3 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Jane",
		lastName:  "Smith",
		age:       25,
	}
	fmt.Println(p3.firstName + " " + p3.lastName)
	fmt.Println(p3.age)
}

type Person struct {
	firstName   string
	lastName    string
	age         int
	address     Address
	ContactInfo // Embedded struct like inheritance, we can access its fields directly using Person struct
}

type Address struct {
	city  string
	state string
}

type ContactInfo struct {
	email  string
	mobile string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) increaseAgeByOne() {
	p.age++
}
