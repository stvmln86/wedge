// Package parse implements program parsing functions.
package parse

import (
	"strings"
	"unicode"

	"github.com/stvmln86/wedge/wedge/atoms/atom"
)

// Clean returns a clean lowercase program string.
func Clean(s string) string {
	var rs []rune
	for _, r := range strings.ToLower(s) {
		switch {
		case unicode.IsSpace(r):
			rs = append(rs, ' ')
		default:
			rs = append(rs, r)
		}
	}

	return strings.TrimSpace(string(rs))
}

// Explode returns a split token slice from a program string.
func Explode(s string) []string {
	return strings.Fields(s)
}

// ParseSlice returns an Atom slice from a token slice.
func ParseSlice(ss []string) ([]atom.Atom, error) {
	var as []atom.Atom
	for _, s := range ss {
		a, err := atom.Atomise(s)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

// ParseString returns an Atom slice from a program string.
func ParseString(s string) ([]atom.Atom, error) {
	return ParseSlice(Explode(Clean(s)))
}
