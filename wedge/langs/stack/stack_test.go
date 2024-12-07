package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/wedge/wedge/atoms/cell"
)

func TestNew(t *testing.T) {
	// success
	s := New(1, 2)
	assert.Equal(t, []cell.Cell{1, 2}, s.Cells)
}

func TestEmpty(t *testing.T) {
	// success - true
	b := New().Empty()
	assert.True(t, b)

	// success - false
	b = New(1, 2).Empty()
	assert.False(t, b)
}

func TestLen(t *testing.T) {
	// success
	i := New(1, 2).Len()
	assert.Equal(t, 2, i)
}

func TestPop(t *testing.T) {
	// setup
	s := New(1)

	// success
	c, err := s.Pop()
	assert.Equal(t, cell.Cell(1), c)
	assert.NoError(t, err)

	// error - Stack is empty
	c, err = s.Pop()
	assert.Zero(t, c)
	assert.EqualError(t, err, "Stack is empty")
}

func TestPopN(t *testing.T) {
	// setup
	s := New(1, 2)

	// success
	cs, err := s.PopN(2)
	assert.Equal(t, []cell.Cell{2, 1}, cs)
	assert.NoError(t, err)

	// error - Stack is missing Cells
	cs, err = s.PopN(1)
	assert.Nil(t, cs)
	assert.EqualError(t, err, "Stack is missing 1 Cells")
}

func TestPush(t *testing.T) {
	// setup
	s := New()

	// success
	s.Push(1)
	assert.Equal(t, []cell.Cell{1}, s.Cells)
}

func TestPushAll(t *testing.T) {
	// setup
	s := New()

	// success
	s.PushAll([]cell.Cell{1, 2})
	assert.Equal(t, []cell.Cell{1, 2}, s.Cells)
}

func TestString(t *testing.T) {
	// success
	s := New(1, 2).String()
	assert.Equal(t, "1 2", s)
}
