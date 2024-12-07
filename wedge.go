///////////////////////////////////////////////////////////////////////////////////////
//             wedge · a minimal stack language in Go · by Stephen Malone            //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                                part 0 · the globals                               //
///////////////////////////////////////////////////////////////////////////////////////

// Stdin is the global input stream.
var Stdin io.Reader = os.Stdin

// Stdout is the global output stream.
var Stdout io.Writer = os.Stdout

///////////////////////////////////////////////////////////////////////////////////////
//                                part 1 · the helpers                               //
///////////////////////////////////////////////////////////////////////////////////////

// Read returns an integer from Stdin.
func Read() int {
	var bs = make([]byte, 1)
	Stdin.Read(bs)
	return int(bs[0])
}

// Write writes an integer or byte to Stdout.
func Write(a any) {
	switch a := a.(type) {
	case byte:
		Stdout.Write([]byte{a})
	case int:
		s := fmt.Sprintf("%c", a)
		Stdout.Write([]byte(s))
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                                 part 2 · the stack                                //
///////////////////////////////////////////////////////////////////////////////////////

// Stack is a last-in-first-out stack of stored integers.
var Stack []int

// Pop removes and returns the top N integers on the Stack.
func Pop(n int) ([]int, error) {
	if len(Stack) < n {
		return nil, fmt.Errorf("Stack is insufficient")
	}

	is := Stack[len(Stack)-n:]
	Stack = Stack[:len(Stack)-n]
	slices.Reverse(is)
	return is, nil
}

// Push appends one or more integers to the top of the Stack.
func Push(is ...int) {
	Stack = append(Stack, is...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                                 part 3 · the queue                                //
///////////////////////////////////////////////////////////////////////////////////////

// Queue is a first-in-first-out queue of parsed atoms.
var Queue []any

// Dequeue removes and returns the next atom in the Queue.
func Dequeue() (any, error) {
	if len(Queue) == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	a := Queue[0]
	Queue = Queue[1:]
	return a, nil
}

// DequeueTo removes and all atoms up to an atom in the Queue.
func DequeueTo(a any) ([]any, error) {
	i := slices.Index(Queue, a)
	if i < 0 {
		return nil, fmt.Errorf("Queue is missing Atom %q", a)
	}

	as := Queue[:i]
	Queue = Queue[i+1:]
	return as, nil
}

// Enqueue appends an atom slice to the end of the Queue.
func Enqueue(as []any) {
	Queue = append(Queue, as...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                                part 4 · the parser                                //
///////////////////////////////////////////////////////////////////////////////////////

// Parse returns a parsed atom slice from a string.
func Parse(s string) []any {
	var as []any

	for _, s := range strings.Fields(strings.ToLower(s)) {
		i, err := strconv.Atoi(s)
		if err == nil {
			as = append(as, i)
		} else {
			as = append(as, s)
		}
	}

	return as
}

///////////////////////////////////////////////////////////////////////////////////////
//                               part 5 · the operators                              //
///////////////////////////////////////////////////////////////////////////////////////

// Oper is a global operator function.
type Oper func() error

// Opers is a global map of all existing Oper functions.
var Opers = make(map[string]Oper)

// Wrap wraps an integer slice function with a Stack error.
func Wrap(n int, ifun func([]int)) Oper {
	return func() error {
		is, err := Pop(n)
		if err != nil {
			return err
		}

		ifun(is)
		return nil
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                                part 6 · the runtime                               //
///////////////////////////////////////////////////////////////////////////////////////

// Evaluate evaluates the next atom in the Queue.
func Evaluate(a any) error {
	switch a := a.(type) {
	case int:
		Push(a)
		return nil
	case string:
		f, b := Opers[a]
		if !b {
			return fmt.Errorf("invalid name %q", a)
		}

		return f()
	default:
		return fmt.Errorf("invalid atom %q", a)
	}
}

// EvaluateQueue evaluates all atoms in the Queue.
func EvaluateQueue() error {
	for len(Queue) > 0 {
		a, err := Dequeue()
		if err != nil {
			return err
		}

		if err := Evaluate(a); err != nil {
			return err
		}
	}

	return nil
}

// EvaluateSlice evaluates all atoms in an atom slice.
func EvaluateSlice(as []any) error {
	for _, a := range as {
		if err := Evaluate(a); err != nil {
			return err
		}
	}

	return nil
}

// EvaluateString evaluates all atoms in a parsed string.
func EvaluateString(s string) error {
	return EvaluateSlice(Parse(s))
}

///////////////////////////////////////////////////////////////////////////////////////
//                               part 7 · the commands                               //
///////////////////////////////////////////////////////////////////////////////////////

// RunREPL launches a read-eval-print loop.
func RunREPL() {
	r := bufio.NewReader(Stdin)

	for {
		fmt.Fprintf(Stdout, "> ")
		s, _ := r.ReadString('\n')
		Enqueue(Parse(s))

		if err := EvaluateQueue(); err != nil {
			fmt.Fprintf(Stdout, "Error: %s.\n", err.Error())

		} else if len(Stack) > 0 {
			fmt.Fprintf(Stdout, ": %v\n", Stack)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                             part 8 · the main function                            //
///////////////////////////////////////////////////////////////////////////////////////

func init() {
	// Mathematical functions.
	Opers["+"] = Wrap(2, func(is []int) { Push(is[0] + is[1]) })
	Opers["-"] = Wrap(2, func(is []int) { Push(is[0] - is[1]) })
	Opers["*"] = Wrap(2, func(is []int) { Push(is[0] * is[1]) })
	Opers["/"] = Wrap(2, func(is []int) { Push(is[0] / is[1]) })
	Opers["%"] = Wrap(2, func(is []int) { Push(is[0] % is[1]) })

	// Stack functions.
	Opers["&"] = Wrap(1, func(is []int) { Push(is[0], is[0]) })
	Opers["~"] = Wrap(2, func(is []int) { Push(is[0], is[1]) })
	Opers["@"] = Wrap(3, func(is []int) { Push(is[1], is[0], is[2]) })

	// Input/output functions.
	Opers["."] = Wrap(1, func(is []int) { Write(is[0]) })
	Opers[","] = Wrap(0, func(is []int) { Push(Read()) })

	// Logic functions.
	Opers["?["] = func() error {
		is, err := Pop(1)
		if err != nil {
			return err
		}

		as, err := DequeueTo("]?")
		if err != nil {
			return err
		}

		if is[0] != 0 {
			return EvaluateSlice(as)
		}

		return nil
	}
}

func main() {
	// ss := os.Args[1:]
	switch {
	default:
		RunREPL()
	}
}
