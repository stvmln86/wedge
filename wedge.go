package main

import (
	"bufio"
	"fmt"

	"github.com/stvmln86/wedge/wedge"
)

// RunREPL launches a read-eval-print loop.
func RunREPL() {
	r := bufio.NewReader(wedge.Stdin)

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(wedge.Stdout, "Error: %s.\n", r)
		}
	}()

	for {
		fmt.Fprintf(wedge.Stdout, "> ")
		s, _ := r.ReadString('\n')

		wedge.Enqueue(wedge.Parse(s)...)
		wedge.EvaluateQueue()

		if len(wedge.Stack) > 0 {
			fmt.Fprintf(wedge.Stdout, ": %v\n", wedge.Stack)
		}
	}
}

func main() {
	switch {
	default:
		RunREPL()
	}
}
