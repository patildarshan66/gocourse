package intermediate

import (
	"errors"
	"fmt"
)

func main() {

	result, err := square(5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	println("Result:", result)

	result1, err1 := square(0)
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}
	println("Result:", result1)

	data := []byte{}

	if err2 := process(data); err2 != nil {
	}
	err2 := process(data)
	if err2 != nil {
		fmt.Println("Error", err2)
		return
	}
	fmt.Println("Data Processed")

	if err3 := eProcess(); err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("Process Completed")

}

func square(x int) (int, error) {
	if x == 0 {
		return 0, errors.New("Number should not be zero to calculate square root")
	}

	res := x * x
	return res, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Data should not be null")
	}
	return nil
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func eProcess() error {
	return &myError{message: "Custom Error message"}
}
