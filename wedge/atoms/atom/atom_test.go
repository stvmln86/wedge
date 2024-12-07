package atom

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/wedge/wedge/atoms/cell"
	"github.com/stvmln86/wedge/wedge/atoms/word"
)

func TestAtomise(t *testing.T) {
	// success - Cell
	a, err := Atomise("123")
	assert.Equal(t, cell.Cell(123), a)
	assert.NoError(t, err)

	// success - Word
	a, err = Atomise("word")
	assert.Equal(t, word.Word("word"), a)
	assert.NoError(t, err)

	// error - invalid Atom
	a, err = Atomise("")
	assert.Nil(t, a)
	assert.EqualError(t, err, `invalid Atom ""`)
}
