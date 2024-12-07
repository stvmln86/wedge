package wedge

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/wedge/wedge/atoms/cell"
)

func TestAdd2(t *testing.T) {
	// setup
	w := mockWedge("")
	w.Stack.Cells = []cell.Cell{1, 2}

	// success
	err := Add2(w)
	assert.Equal(t, []cell.Cell{3}, w.Stack.Cells)
	assert.NoError(t, err)
}
