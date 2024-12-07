package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeue(t *testing.T) {
	// setup
	Queue = []any{"a"}

	// success
	a := Dequeue()
	assert.Equal(t, "a", a)
	assert.Empty(t, Queue)
}

func TestDequeueTo(t *testing.T) {
	// setup
	Queue = []any{"a", "end"}

	// success
	as := DequeueTo("end")
	assert.Equal(t, []any{"a"}, as)
	assert.Empty(t, Queue)
}

func TestEnqueue(t *testing.T) {
	// setup
	Queue = []any{}

	// success
	Enqueue("a")
	assert.Equal(t, []any{"a"}, Queue)
}
