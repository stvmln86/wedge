///////////////////////////////////////////////////////////////////////////////////////
//             wedge · a minimal stack language in Go · by Stephen Malone            //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"slices"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                            part 1 · globals & constants                           //
///////////////////////////////////////////////////////////////////////////////////////

// Opers is the global operator function map.
var Opers = make(map[string]func())

// Running is the global program running indicator.
var Running = true

// Stdin is the global input stream.
var Stdin io.Reader = os.Stdin

// Stdout is the global output stream.
var Stdout io.Writer = os.Stdout

///////////////////////////////////////////////////////////////////////////////////////
//                              part 2 · input & output                              //
///////////////////////////////////////////////////////////////////////////////////////

// Read returns an input byte from Stdin as an integer.
func Read() int {
	var bs = make([]byte, 1)
	Stdin.Read(bs)
	return int(bs[0])
}

// Write writes a byte or integer to Stdout.
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
//                           part 3 · the stack & functions                          //
///////////////////////////////////////////////////////////////////////////////////////

// Stack is a last-in-first-out stack of stored integers.
var Stack []int

// Peek returns the top integer on the Stack.
func Peek() int {
	if len(Stack) == 0 {
		panic("stack is insufficient")
	}

	return Stack[len(Stack)-1]
}

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

///////////////////////////////////////////////////////////////////////////////////////
//                           part 4 · the queue & functions                          //
///////////////////////////////////////////////////////////////////////////////////////

// Queue is a first-in-first-out queue of parsed atoms.
var Queue []any

// Dequeue removes and returns the next atom in the Queue.
func Dequeue() any {
	if len(Queue) == 0 {
		panic("queue is empty")
	}

	a := Queue[0]
	Queue = Queue[1:]
	return a
}

// DequeueTo removes and returns all atoms up to an atom in the Queue.
func DequeueTo(a any) []any {
	i := slices.Index(Queue, a)
	if i < 0 {
		panic(fmt.Sprintf("queue is missing %q", a))
	}

	as := Queue[:i]
	Queue = Queue[i+1:]
	return as
}

// Enqueue appends an atom slice to the end of the Queue.
func Enqueue(as []any) {
	Queue = append(Queue, as...)
}

// Insert inserts an atom slice at the start of the Queue.
func Insert(as []any) {
	Queue = append(as, Queue...)
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part 5 · parsing & evaluating                           //
///////////////////////////////////////////////////////////////////////////////////////

// Parse returns a parsed atom slice from a string.
func Parse(s string) []any {
	var as []any

	for _, s := range strings.Fields(strings.ToLower(s)) {
		if i, err := strconv.Atoi(s); err == nil {
			as = append(as, i)
		} else {
			as = append(as, s)
		}
	}

	return as
}

// Evaluate evaluates the next atom in the Queue.
func Evaluate(a any) {
	switch a := a.(type) {
	case int:
		Push(a)
	case string:
		f, ok := Opers[a]
		if !ok {
			panic(fmt.Sprintf("invalid reference %q", a))
		}

		f()
	default:
		panic(fmt.Sprintf("invalid atom type %T", a))
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part 6 · operator functions                            //
///////////////////////////////////////////////////////////////////////////////////////

// InitOpers initialises the Opers map with default operator functions.
func InitOpers() {
	// Mathematical operators.
	Opers["+"] = func() { Push(Pop() + Pop()) }
	Opers["-"] = func() { Push(Pop() - Pop()) }
	Opers["*"] = func() { Push(Pop() * Pop()) }
	Opers["/"] = func() { Push(Pop() / Pop()) }
	Opers["%"] = func() { Push(Pop() % Pop()) }

	// Stack operators.
	Opers["&"] = func() { Push(Peek()) }
	Opers["#"] = func() { Push(len(Stack)) }
	Opers["~"] = func() { Push(Pop(), Pop()) }
	Opers["@"] = func() {
		a, b, c := Pop(), Pop(), Pop()
		Push(b, a, c)
	}

	// Input/output operators.
	Opers["."] = func() { Write(Pop()) }
	Opers[","] = func() { Push(Read()) }

	// Logic operators.
	Opers["{?"] = func() {
		as := DequeueTo("?}")
		if Pop() != 0 {
			Insert(as)
		}
	}

	Opers["{#"] = func() {
		as := DequeueTo("#}")
		for range Pop() {
			Insert(as)
		}
	}

	Opers["{="] = func() {
		as := DequeueTo("=}")
		if _, ok := as[0].(string); len(as) < 2 || !ok {
			panic("invalid function definition")
		}

		Opers[as[0].(string)] = func() { Insert(as[1:]) }
	}

	// System operators.
	Opers["·"] = func() {}
	Opers["dump"] = func() { fmt.Fprintf(Stdout, ": %v\n", Stack) }
	Opers["exit"] = func() { Running = false }
	Opers["eval"] = func() {
		var rs []rune
		for len(Stack) > 0 {
			rs = append(rs, rune(Pop()))
		}

		Insert(Parse(string(rs)))
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part 7 · top-level commands                            //
///////////////////////////////////////////////////////////////////////////////////////

// RunREPL launches a read-eval-print loop.
func RunREPL() {
	r := bufio.NewReader(Stdin)

	for Running {
		fmt.Fprintf(Stdout, "> ")
		s, _ := r.ReadString('\n')
		as := Parse(s)

		if len(as) != 0 {
			Enqueue(as)
			Enqueue(Parse("# {? dump ?}"))
			for len(Queue) > 0 {
				Evaluate(Dequeue())
			}
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                              part 8 · main functions                              //
///////////////////////////////////////////////////////////////////////////////////////

func init() {
	InitOpers()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(Stdout, "Error: %s.\n", r)
		}
	}()

	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt)
		<-s
		os.Exit(0)
	}()

	RunREPL()
}
