package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/wedge/wedge/atoms/atom"
	"github.com/stvmln86/wedge/wedge/atoms/cell"
	"github.com/stvmln86/wedge/wedge/atoms/word"
)

func TestClean(t *testing.T) {
	// success
	s := Clean("\t123 ABC\n")
	assert.Equal(t, "123 abc", s)
}

func TestExplode(t *testing.T) {
	// success
	ss := Explode("  123  abc  ")
	assert.Equal(t, []string{"123", "abc"}, ss)
}

func TestParseSlice(t *testing.T) {
	// success
	as, err := ParseSlice([]string{"123", "abc"})
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, as)
	assert.NoError(t, err)
}

func TestParseString(t *testing.T) {
	// success
	as, err := ParseString("\t123 ABC\n")
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, as)
	assert.NoError(t, err)
}
