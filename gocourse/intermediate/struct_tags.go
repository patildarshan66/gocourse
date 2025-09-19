package intermediate

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`

	// FirstName string `json:"firstName" db:"firstName" xml:"firstName"` // example of multiple tags

	// Age       int    `json:"age,omitempty"` //used omitempty to ignore this key when value is nil, json marshal

	//Age       int    `json:"-"` // used - to ignore this key in json marshal
}

func main() {

	person := Person{FirstName: "Darshan", LastName: "Patil", Age: 26}

	jsonData, err := json.Marshal(person)

	if err != nil {
		log.Fatalln("Error while marshaling person:", err)
	}

	println("Json Data:", string(jsonData))

}
