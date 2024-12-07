// Package stack implements the Stack type.
package stack

import (
	"fmt"
	"slices"
	"strings"

	"github.com/stvmln86/wedge/wedge/atoms/cell"
)

// Stack is a last-in-first-out stack of Cells.
type Stack struct {
	Cells []cell.Cell
}

// New returns a new Stack from zero or more Cells.
func New(cs ...cell.Cell) *Stack {
	return &Stack{cs}
}

// Empty returns true if the Stack has no Cells.
func (s *Stack) Empty() bool {
	return len(s.Cells) == 0
}

// Len returns the number of Cells on the Stack.
func (s *Stack) Len() int {
	return len(s.Cells)
}

// Pop removes and returns the top Cell on the Stack.
func (s *Stack) Pop() (cell.Cell, error) {
	if len(s.Cells) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}

	c := s.Cells[len(s.Cells)-1]
	s.Cells = s.Cells[:len(s.Cells)-1]
	return c, nil
}

// PopN removes and returns the top N Cells on the Stack.
func (s *Stack) PopN(n int) ([]cell.Cell, error) {
	if len(s.Cells) < n {
		return nil, fmt.Errorf("Stack is missing %d Cells", n)
	}

	cs := s.Cells[len(s.Cells)-n:]
	s.Cells = s.Cells[:len(s.Cells)-n]
	slices.Reverse(cs)
	return cs, nil
}

// Push appends a Cell to the top of the Stack.
func (s *Stack) Push(c cell.Cell) {
	s.Cells = append(s.Cells, c)
}

// PushAll appends a Cell slice to the top of the Stack.
func (s *Stack) PushAll(cs []cell.Cell) {
	s.Cells = append(s.Cells, cs...)
}

// String returns the Stack as a string.
func (s *Stack) String() string {
	var ss []string
	for _, c := range s.Cells {
		ss = append(ss, c.String())
	}

	return strings.Join(ss, " ")
}
