package main

// Stack is a last-in-first-out stack of stored integers.
var Stack []int

// Pop removes and returns the top integer on the Stack.
func Pop() int {
	if len(Stack) == 0 {
		panic("stack is insufficient")
	}

	i := Stack[len(Stack)-1]
	Stack = Stack[:len(Stack)-1]
	return i
}

// Push appends one or more integers to the top of the Stack.
func Push(is ...int) {
	Stack = append(Stack, is...)
}