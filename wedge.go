///////////////////////////////////////////////////////////////////////////////////////
//             wedge · a minimal stack language in Go · by Stephen Malone            //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

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
type Oper func()

// Opers is a global map of all existing Oper functions.
var Opers = make(map[string]Oper)

///////////////////////////////////////////////////////////////////////////////////////
//                                part 6 · the runtime                               //
///////////////////////////////////////////////////////////////////////////////////////

// Evaluate evaluates the next atom in the Queue.
func Evaluate(a any) {
	switch a := a.(type) {
	case int:
		Push(a)
	case string:
		f, b := Opers[a]
		if !b {
			panic("invalid reference")
		}

		f()
	default:
		panic("invalid atom type")
	}
}

// EvaluateQueue evaluates all atoms in the Queue.
func EvaluateQueue() {
	for len(Queue) > 0 {
		Evaluate(Dequeue())
	}
}

// EvaluateSlice evaluates all atoms in an atom slice.
func EvaluateSlice(as []any) {
	for _, a := range as {
		Evaluate(a)
	}
}

// EvaluateString evaluates all atoms in a parsed string.
func EvaluateString(s string) {
	EvaluateSlice(Parse(s))
}

///////////////////////////////////////////////////////////////////////////////////////
//                               part 7 · the commands                               //
///////////////////////////////////////////////////////////////////////////////////////

// RunREPL launches a read-eval-print loop.
func RunREPL() {
	r := bufio.NewReader(Stdin)

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(Stdout, "Error: %s.\n", r)
		}
	}()

	for {
		fmt.Fprintf(Stdout, "> ")
		s, _ := r.ReadString('\n')
		Enqueue(Parse(s))

		EvaluateQueue()
		if len(Stack) > 0 {
			fmt.Fprintf(Stdout, ": %v\n", Stack)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                             part 8 · the main function                            //
///////////////////////////////////////////////////////////////////////////////////////

func init() {
	// Mathematical functions.
	Opers["+"] = func() { Push(Pop() + Pop()) }
	Opers["-"] = func() { Push(Pop() - Pop()) }
	Opers["*"] = func() { Push(Pop() * Pop()) }
	Opers["/"] = func() { Push(Pop() / Pop()) }
	Opers["%"] = func() { Push(Pop() % Pop()) }

	// Stack functions.
	Opers["&"] = func() { Push(Stack[len(Stack)-1]) }
	// Opers["~"] = func() { Push(is[0], is[1]) }
	// Opers["@"] = func() { Push(is[1], is[0], is[2]) }
	Opers["#"] = func() { Push(len(Stack)) }

	// Input/output functions.
	Opers["."] = func() { Write(Pop()) }
	Opers[","] = func() { Push(Read()) }

	// Logic functions.
	Opers["?["] = func() {
		as := DequeueTo("]?")
		if Pop() != 0 {
			EvaluateSlice(as)
		}
	}
}

func main() {
	// ss := os.Args[1:]
	switch {
	default:
		RunREPL()
	}
}
