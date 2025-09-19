package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	person := Person{Name: "Darshan", Email: "text@gmail.com"}

	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error Marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{Name: "Kalpesh", Age: 27, Email: "test@gmail.co.in", Address: Address{City: "Nandurbar", Pincode: "435432"}}

	jsonData1, err := json.Marshal(person1)

	if err != nil {
		fmt.Println("Error Marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData1))

	jsonStr := `{"name":"Kalpesh","age":27,"email":"test@gmail.co.in","Address":{"city":"Nandurbar","pincode":"435432"}}`
	var personData Person
	err = json.Unmarshal([]byte(jsonStr), &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling to JSON:", err)
		return
	}
	fmt.Println(personData)

	listOfAddress := []Address{
		{City: "Nandurbar", Pincode: "123456"},
		{City: "Shahada", Pincode: "435209"},
		{City: "Mhasawad", Pincode: "425432"},
	}

	data1, err := json.Marshal(listOfAddress)
	if err != nil {
		fmt.Println("Error Unmarshalling to JSON:", err)
		return
	}

	fmt.Println(string(data1))

	//Handling Unknown JSON  Structure

	jsonData2 := `{"name":"Kalpesh","age":27,"email":"test@gmail.co.in","Address":{"city":"Nandurbar","pincode":"435432"}}`
	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		log.Fatalln("Error while unmarshalling:", err)
	}
	fmt.Println("Decoded Json:", data)
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email"`
	Address Address `json:"Address"`
}

type Address struct {
	City    string `json:"city"`
	Pincode string `json:"pincode"`
}
