package main

import (
	"strconv"
	"strings"
)

// Parse returns a parsed atom slice from a string.
func Parse(s string) []any {
	var as []any

	for _, s := range strings.Fields(strings.ToLower(s)) {
		if i, err := strconv.Atoi(s); err == nil {
			as = append(as, i)
		} else {
			as = append(as, s)
		}
	}

	return as
}
