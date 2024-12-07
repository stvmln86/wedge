package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stvmln86/wedge/wedge/langs/wedge"
)

func main() {
	w := wedge.New(os.Stdin, os.Stdout)
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		s, _ := r.ReadString('\n')

		if err := w.EvaluateString(s); err != nil {
			fmt.Printf("Error: %s.\n", err.Error())

		} else if !w.Stack.Empty() {
			fmt.Printf(": %s\n", w.Stack)
		}
	}
}
