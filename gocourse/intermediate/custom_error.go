package intermediate

import (
	"errors"
	"fmt"
)

func main() {
	// err := doSomething()
	data := []byte{}
	data = append(data, 10)
	result, err := processFileBytes(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Operation completed successfully.", result)

}

type customError struct {
	statusCode int
	message    string
	err        error
}

// Error returns the error message. Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v", e.statusCode, e.message, e.err)
}

// func doSomething() error {
// 	return &customError{statusCode: 500, message: "Something went wrong!"}
// }

func processFileBytes(data []byte) (bool, error) {
	result, err := saveFile(data)
	if err != nil {
		return result, &customError{statusCode: 500, message: "Something went wrong!", err: err}
	}
	return result, nil
}

func saveFile(data []byte) (bool, error) {
	if len(data) == 0 {
		return false, errors.New("File data is empty. Unable to save file")
	}
	return true, nil
}
