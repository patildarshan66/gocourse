package basics

import "fmt"

func main() {
	process(10)
	fmt.Println(" This statement will be executed after process function")
	process(-1)
	fmt.Println(" This statement will not be executed after process function beacause of panic")
	processWithRecover()
	fmt.Println(" This statement will be executed after process function beacause of panic recovery")

}

func process(processId int) {
	defer fmt.Println("defer 1:=> Means It will be executed at the end of the function")
	defer fmt.Println("defer 2:=> In case of panice All Defer will be executed")
	defer fmt.Println("defer 3:=> Defer will be executed in Last In First Out Order")
	if processId < 0 {
		panic("Panic :=> Process Id can not be negative and stop the execution of the function . Like Throuwing an Error, Afrer Panice nothing will be excuted")
	}
	fmt.Println("Process Id is: ", processId)
}

func processWithRecover() {
	defer func() {
		// if r := recover(); r != nil {
		r := recover() // recover function is used to catch the panic and continue the execution of the program,
		// recover will return nil if there is no panic
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Process start With Recover Function")
	panic("Panic :=> Something went wrong")
	fmt.Println("Process end With Recover Function")
}
