// Package atom implements the Atom interface and functions.
package atom

import (
	"fmt"

	"github.com/stvmln86/wedge/wedge/atoms/cell"
)

// Atom is a parsed program value.
type Atom interface {
	// Bool returns the Atom as a boolean.
	Bool() bool

	// Native returns the Atom as a native value.
	Native() any

	// String returns the Atom as a string.
	String() string
}

// Atomise returns an Atom from a string.
func Atomise(s string) (Atom, error) {
	switch {
	case cell.Is(s):
		return cell.Parse(s)
	default:
		return nil, fmt.Errorf("invalid Atom %q", s)
	}
}
