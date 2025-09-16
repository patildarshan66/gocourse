package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

}

func main() {
	tempFileName := "tempFileDemo"
	tempFile, err := os.CreateTemp(".", tempFileName)
	checkError(err)
	fmt.Println("Temp File is created successfully:", tempFile.Name())
	defer checkError(os.Remove(tempFile.Name()))
	defer checkError(tempFile.Close())

	tempDir, err := os.MkdirTemp(".", tempFileName)
	checkError(err)

	defer os.Remove(tempDir)

	fmt.Println("Temporary directory created:", tempDir)

}
