package atom

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/wedge/wedge/atoms/cell"
)

func TestAtomise(t *testing.T) {
	// success - Cell
	a, err := Atomise("123")
	assert.Equal(t, cell.Cell(123), a)
	assert.NoError(t, err)

	// error - invalid Atom
	a, err = Atomise("nope")
	assert.Nil(t, a)
	assert.EqualError(t, err, `invalid Atom "nope"`)
}
