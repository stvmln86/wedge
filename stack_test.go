package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	// setup
	Stack = []int{1}

	// success
	i := Pop()
	assert.Equal(t, 1, i)
	assert.Empty(t, Stack)
}

func TestPush(t *testing.T) {
	// setup
	Stack = []int{}

	// success
	Push(1)
	assert.Equal(t, []int{1}, Stack)
}
