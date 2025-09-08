package intermediate

import "fmt"

func main() {

	a := 1
	b := 2

	firstName := "Darshan"
	lastName := "Patil"

	a, b = swap(a, b)
	firstName, lastName = swap(firstName, lastName)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(firstName)
	fmt.Println(lastName)

	intStack := stack[int]{}

	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printStack()
	fmt.Println(intStack.pop())
	fmt.Println("Is stack is empty:", intStack.isEmpty())
	intStack.printStack()
	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack is empty:", intStack.isEmpty())

	stringStack := stack[string]{}

	stringStack.push("Darshan")
	stringStack.push("Kalpesh")
	stringStack.push("Rohit")
	stringStack.printStack()
	fmt.Println(stringStack.pop())
	fmt.Println("Is stack is empty:", stringStack.isEmpty())
	stringStack.printStack()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stack is empty:", stringStack.isEmpty())
}

func swap[T any](a, b T) (T, T) {
	return b, a
}

type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(t T) {
	s.elements = append(s.elements, t)
}

func (s *stack[T]) pop() (T, bool) {
	if s.isEmpty() {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:(len(s.elements) - 1)]
	return element, true
}

func (s stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s stack[T]) printStack() {
	for _, value := range s.elements {
		fmt.Print(value)
	}
	fmt.Println()
}
