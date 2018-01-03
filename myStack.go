/*
myStack: set of functions that implement a stack including the following four
methods:
a. Push: Adds a data element to the top of the stack
b. Pop: Removes a data element from the top of the stack
c. Size: Returns the total number of elements in the stack
d. isEmpty: Returns true if the stack is empty.
To run: go run myStack.go
*/
package main

import (
	"fmt"
)

type Stack struct {
	top    *Element
	length int
}

type Element struct {
	value interface{} // the empty interface.
	next  *Element
}

func main() {
	st := Stack{} // create an empty stack
	st.Push(5)
	st.Push(4)
	st.Push(3)
	fmt.Printf("stack before push: %v\n", st)
	st.Push(1) // push 1 to the top of stack
	fmt.Printf("stack after push: %v\n", st)
	fmt.Printf("pop the top of stack: %d and stack after pop: %v\n", st.Pop(), st)
	fmt.Printf("Size of the stack: %d\n", st.Size())
	fmt.Printf("Is the stack empty? %t\n", st.IsEmpty())

	for !st.IsEmpty() {
		fmt.Printf("pop the top of stack: %d\n", st.Pop())
	}
	fmt.Printf("Is the stack empty? %t\n", st.IsEmpty())
}

// Return the stack's length
func (s *Stack) Size() int {
	return s.length
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.length++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.length > 0 {
		value, s.top = s.top.value, s.top.next
		s.length--
		return
	}
	return nil
}

// Returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	if s.length == 0 {
		return true
	}
	return false
}
