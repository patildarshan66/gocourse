package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(strings.NewReader("I am darshan patil. My age is 26.\n"))

	// data := make([]byte, 20) // it will read only first 20 bytes of input string
	// n, err := reader.Read(data)

	// if err != nil {
	// 	fmt.Println("Error while reading:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// readStr, err := reader.ReadString('\n') // stop reading when I will get \n in input string
	// //It will start reading where it was last stoped ex. data := make([]byte, 20) it was read 20 bytes already then start reading from 21st byte

	// if err != nil {
	// 	fmt.Println("Error while reading:", err)
	// 	return
	// }
	// fmt.Println("Read String:", readStr)

	// Stdout Writer with writing a slice
	writer1 := bufio.NewWriter(os.Stdout)

	data1 := []byte("Hello, My Self Darshan. My Age is 25.\n")

	n, err := writer1.Write(data1)

	if err != nil {
		fmt.Println("Error while writing to buffer:", err)
	}
	fmt.Printf("Wrote %d bytes\n", n)
	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer1.Flush()

	if err != nil {
		fmt.Println("Error while flusing the buffer:", err)
	}
	///My Custom Writer with writing a slice
	r := MyCustomWriter{}
	writer := bufio.NewWriter(&r)

	data := []byte("Hello, My Self Darshan!")

	n1, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error while writing to buffer:", err)
	}
	fmt.Printf("Wrote %d bytes\n", n1)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error while flusing the buffer:", err)
	}

	//Writing a string using My Custom
	str := "This is a string\n"
	n2, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error while writing to buffer:", err)
	}
	fmt.Printf("Wrote %d bytes\n", n2)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error while flusing the buffer:", err)
	}

	//Writing a string using Stdout
	str1 := "This is a string\n"
	n3, err := writer.WriteString(str1)
	if err != nil {
		fmt.Println("Error while writing to buffer:", err)
	}
	fmt.Printf("Wrote %d bytes\n", n3)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error while flusing the buffer:", err)
	}

}

type MyCustomWriter struct {
}

func (myWriter *MyCustomWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(string(p)), nil
}

func (myWriter *MyCustomWriter) WriteString(p string) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
