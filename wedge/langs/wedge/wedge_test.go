package wedge

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/wedge/wedge/atoms/atom"
	"github.com/stvmln86/wedge/wedge/atoms/cell"
	"github.com/stvmln86/wedge/wedge/atoms/word"
)

func mockWedge(s string) *Wedge {
	r := bytes.NewBufferString(s)
	w := bytes.NewBuffer(nil)
	return New(r, w)
}

func TestNew(t *testing.T) {
	// success
	w := mockWedge("")
	assert.NotNil(t, w)
}

func TestEvaluate(t *testing.T) {
	// setup
	w := mockWedge("")
	w.Stack.Cells = []cell.Cell{1}
	w.Queue.Atoms = []atom.Atom{cell.Cell(2), word.Word("+")}

	// success - Cell
	err := w.Evaluate()
	assert.Equal(t, []cell.Cell{1, 2}, w.Stack.Cells)
	assert.NoError(t, err)

	// success - Word
	err = w.Evaluate()
	assert.Equal(t, []cell.Cell{3}, w.Stack.Cells)
	assert.NoError(t, err)
}

func TestEvaluateAll(t *testing.T) {
	// setup
	w := mockWedge("")
	w.Queue.Atoms = []atom.Atom{cell.Cell(1), cell.Cell(2), word.Word("+")}

	// success
	err := w.EvaluateAll()
	assert.Equal(t, []cell.Cell{3}, w.Stack.Cells)
	assert.NoError(t, err)
}

func TestEvaluateString(t *testing.T) {
	// setup
	w := mockWedge("")

	// success
	err := w.EvaluateString("1 2 +")
	assert.Equal(t, []cell.Cell{3}, w.Stack.Cells)
	assert.NoError(t, err)
}
