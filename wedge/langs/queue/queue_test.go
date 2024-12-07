package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/wedge/wedge/atoms/atom"
	"github.com/stvmln86/wedge/wedge/atoms/cell"
	"github.com/stvmln86/wedge/wedge/atoms/word"
)

func TestNew(t *testing.T) {
	// success
	q := New(cell.Cell(123), word.Word("abc"))
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, q.Atoms)
}

func TestDequeue(t *testing.T) {
	// setup
	q := New(cell.Cell(123))

	// success
	a, err := q.Dequeue()
	assert.Equal(t, cell.Cell(123), a)
	assert.NoError(t, err)

	// error - Queue is empty
	a, err = q.Dequeue()
	assert.Nil(t, a)
	assert.EqualError(t, err, "Queue is empty")
}

func TestDequeueTo(t *testing.T) {
	// setup
	q := New(cell.Cell(123), word.Word("abc"))

	// success
	as, err := q.DequeueTo(word.Word("abc"))
	assert.Equal(t, []atom.Atom{cell.Cell(123)}, as)
	assert.NoError(t, err)

	// error - Queue is empty
	as, err = q.DequeueTo(word.Word("nope"))
	assert.Nil(t, as)
	assert.EqualError(t, err, `Queue is missing Atom "nope"`)
}

func TestEmpty(t *testing.T) {
	// success - true
	b := New().Empty()
	assert.True(t, b)

	// success - false
	b = New(cell.Cell(123), word.Word("abc")).Empty()
	assert.False(t, b)
}

func TestEnqueue(t *testing.T) {
	// setup
	q := New()

	// success
	q.Enqueue(cell.Cell(123))
	assert.Equal(t, []atom.Atom{cell.Cell(123)}, q.Atoms)
}

func TestEnqueueAll(t *testing.T) {
	// setup
	q := New()

	// success
	q.EnqueueAll([]atom.Atom{cell.Cell(123), word.Word("abc")})
	assert.Equal(t, []atom.Atom{cell.Cell(123), word.Word("abc")}, q.Atoms)
}

func TestLen(t *testing.T) {
	// success
	i := New(cell.Cell(123), word.Word("abc")).Len()
	assert.Equal(t, 2, i)
}
