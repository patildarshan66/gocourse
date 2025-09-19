package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email,omitempty"`
	Address Address  `xml:"Address"`
}

type Address struct {
	City    string `xml:"city"`
	Pincode string `xml:"pincode"`
}

func handleError(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func main() {

	person := Person{Name: "Darshan", Age: 26, Address: Address{City: "Nandurbar", Pincode: "425432"}}

	xmlData, err := xml.Marshal(person)
	handleError(err)

	println(string(xmlData))
	println("------------------------")

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	handleError(err)

	println(string(xmlData1))
	println("------------------------")
	xmlStr := `<person><name>Darshan</name><age>26</age><email>test@gmail.com</email><Address><city>Nandurbar</city><pincode>425432</pincode></Address></person>`

	var xmlPersonData Person

	err = xml.Unmarshal([]byte(xmlStr), &xmlPersonData)
	handleError(err)
	fmt.Println(xmlPersonData)

	fmt.Println("Local String:", xmlPersonData.XMLName.Local)
	fmt.Println("Namespace:", xmlPersonData.XMLName.Space)
	println("------------------------")
	book := Book{
		ISBN:       "12-45-52-22-34-42",
		Title:      "Go Lang",
		Author:     "Darshan",
		Pseudo:     "Pseudo",
		PseudoAttr: "PseudoAttr",
	}

	bookXml, err := xml.MarshalIndent(book, "", "  ")
	handleError(err)
	println(string(bookXml))
	println("------------------------")

}

type Book struct {
	XMLName    xml.Name
	ISBN       string `xml:"isbn,attr"`
	Title      string `xml:"title,attr"`
	Author     string `xml:"author,attr"`
	Pseudo     string `xml:"isbn"`
	PseudoAttr string `xml:"pseudoAttr,attr"`
}
