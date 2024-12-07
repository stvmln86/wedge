package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	// setup
	Stack = []int{123}

	// success
	i := Pop()
	assert.Equal(t, 123, i)
	assert.Empty(t, Stack)
}

func TestPush(t *testing.T) {
	// setup
	Stack = []int{}

	// success
	Push(123)
	assert.Equal(t, []int{123}, Stack)
}
