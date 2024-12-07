package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	// success - true
	b := Is("123")
	assert.True(t, b)

	// success - false
	b = Is("nope")
	assert.False(t, b)
}

func TestParse(t *testing.T) {
	// success
	c, err := Parse("123")
	assert.Equal(t, Cell(123), c)
	assert.NoError(t, err)

	// error - invalid Cell
	c, err = Parse("nope")
	assert.Zero(t, c)
	assert.EqualError(t, err, `invalid Cell "nope"`)
}

func TestBool(t *testing.T) {
	// success - true positive
	b := Cell(123).Bool()
	assert.True(t, b)

	// success - true negative
	b = Cell(-123).Bool()
	assert.True(t, b)

	// success - false
	b = Cell(0).Bool()
	assert.False(t, b)
}

func TestNative(t *testing.T) {
	// success
	i := Cell(123).Native()
	assert.Equal(t, 123, i)
}

func TestString(t *testing.T) {
	// success
	s := Cell(123).String()
	assert.Equal(t, "123", s)
}
