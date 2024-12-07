package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// success
	as := Parse("\t123 ABC\n")
	assert.Equal(t, []any{123, "abc"}, as)
}
