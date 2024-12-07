// Package cell implements the Cell Atom type.
package cell

import (
	"fmt"
	"strconv"
)

// Cell is a parsed program integer value.
type Cell int

// Is returns true if a string represents a Cell.
func Is(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Parse returns a new Cell from a string.
func Parse(s string) (Cell, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid Cell %q", s)
	}

	return Cell(i), nil
}

// Bool returns the Cell as a boolean.
func (c Cell) Bool() bool {
	return int(c) != 0
}

// Native returns the Cell as a native value.
func (c Cell) Native() any {
	return int(c)
}

// String returns the Cell as a string.
func (c Cell) String() string {
	return strconv.Itoa(int(c))
}
