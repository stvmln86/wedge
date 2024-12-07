package wedge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	// setup
	Stack = []int{}

	// success - int
	Evaluate(123)
	assert.Equal(t, []int{123}, Stack)

	// success - string
	Evaluate("&")
	assert.Equal(t, []int{123, 123}, Stack)
}

func TestEvaluateQueue(t *testing.T) {
	// setup
	Queue = []any{123, "&"}
	Stack = []int{}

	// success
	EvaluateQueue()
	assert.Equal(t, []int{123, 123}, Stack)
}

func TestEvaluateSlice(t *testing.T) {
	// setup
	Stack = []int{}

	// success
	EvaluateSlice([]any{123, "&"})
	assert.Equal(t, []int{123, 123}, Stack)
}

func TestEvaluateString(t *testing.T) {
	// setup
	Stack = []int{}

	// success
	EvaluateString("123 &")
	assert.Equal(t, []int{123, 123}, Stack)
}
