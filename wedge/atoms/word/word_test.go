package word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	// success - true
	b := Is("word")
	assert.True(t, b)

	// success - false
	b = Is("")
	assert.False(t, b)
}

func TestParse(t *testing.T) {
	// success
	w, err := Parse("word\n")
	assert.Equal(t, Word("word"), w)
	assert.NoError(t, err)

	// error - invalid Word
	w, err = Parse("\n")
	assert.Empty(t, w)
	assert.EqualError(t, err, `invalid Word ""`)
}

func TestBool(t *testing.T) {
	// success - true
	b := Word("word").Bool()
	assert.True(t, b)

	// success - false
	b = Word("").Bool()
	assert.False(t, b)
}

func TestNative(t *testing.T) {
	// success
	s := Word("word").Native()
	assert.Equal(t, "word", s)
}

func TestString(t *testing.T) {
	// success
	s := Word("word").String()
	assert.Equal(t, "word", s)
}
