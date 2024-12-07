package main

// Evaluate evaluates the next atom in the Queue.
func Evaluate(a any) {
	switch a := a.(type) {
	case int:
		Push(a)
	case string:
		f, b := Opers[a]
		if !b {
			panic("invalid reference")
		}

		f()
	default:
		panic("invalid atom type")
	}
}

// EvaluateQueue evaluates all atoms in the Queue.
func EvaluateQueue() {
	for len(Queue) > 0 {
		Evaluate(Dequeue())
	}
}

// EvaluateSlice evaluates all atoms in an atom slice.
func EvaluateSlice(as []any) {
	for _, a := range as {
		Evaluate(a)
	}
}

// EvaluateString evaluates all atoms in a parsed string.
func EvaluateString(s string) {
	EvaluateSlice(Parse(s))
}
