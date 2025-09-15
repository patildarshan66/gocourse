package intermediate

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Error while creating a file:", err)
		return
	}

	defer file.Close()

	data := []byte("Hello, I am Darshan Patil\n ")

	// write to file using bytes
	_, err = file.Write(data)

	if err != nil {
		fmt.Println("Error while writing a file:", err)
		return
	}

	fmt.Println("Data has been written to file successfully.")

	file, err = os.Create("WriteString.txt")

	if err != nil {
		fmt.Println("Error while creating a file:", err)
		return
	}

	// write to file using string
	_, err = file.WriteString("Hello, I am learning go lang\n\n")

	if err != nil {
		fmt.Println("Error while writing a file:", err)
		return
	}

	fmt.Println("Data has been written to file successfully using WriteString.")

}
