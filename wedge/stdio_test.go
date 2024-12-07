package wedge

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	// setup
	Stdin = bytes.NewBufferString("t")

	// success
	i := Read()
	assert.Equal(t, int('t'), i)
}

func TestWrite(t *testing.T) {
	// setup
	b := bytes.NewBuffer(nil)
	Stdout = b

	// success
	Write(byte('t'))
	Write(int('t'))
	assert.Equal(t, "tt", b.String())
}
