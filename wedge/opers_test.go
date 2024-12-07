package wedge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// success
	for s, is := range map[string][]int{
		// Mathematical functions.
		"1 2 +": {3},
		"1 2 -": {1},
		"2 3 *": {6},
		"3 6 /": {2},
		"2 5 %": {1},

		// Stack functions.
		"1 &":     {1, 1},
		"1 #":     {1, 1},
		"1 2 ~":   {2, 1},
		"1 2 3 @": {2, 3, 1},
	} {
		Queue = []any{}
		Stack = []int{}

		Enqueue(Parse(s)...)
		EvaluateQueue()
		assert.Equal(t, is, Stack)
	}

	// TODO.
}
