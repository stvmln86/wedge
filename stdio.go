package main

import (
	"fmt"
	"io"
	"os"
)

// Stdin is the global input stream.
var Stdin io.Reader = os.Stdin

// Stdout is the global output stream.
var Stdout io.Writer = os.Stdout

// Read returns an integer from Stdin.
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
