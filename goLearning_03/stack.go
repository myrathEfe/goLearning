package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// düzenlenecek...
func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("Stack is empty")
	}
	lastIndex := len(s.items) - 1
	top := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return top, nil
}

func (s *Stack) printStack() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
	} else {
		fmt.Println(s.items)
	}
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func main() {
	myStack := Stack{}
	myStack.Push(20)
	myStack.Push(30)
	myStack.Push(40)
	myStack.Push(50)
	myStack.printStack()
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()
	myStack.printStack()

}
