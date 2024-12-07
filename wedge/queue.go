package wedge

import "slices"

// Queue is a first-in-first-out queue of parsed atoms.
var Queue []any

// Dequeue removes and returns the next atom in the Queue.
func Dequeue() any {
	if len(Queue) == 0 {
		panic("queue is empty")
	}

	a := Queue[0]
	Queue = Queue[1:]
	return a
}

// DequeueTo removes and all atoms up to an atom in the Queue.
func DequeueTo(a any) []any {
	i := slices.Index(Queue, a)
	if i < 0 {
		panic("queue is insufficient")
	}

	as := Queue[:i]
	Queue = Queue[i+1:]
	return as
}

// Enqueue appends an atom sequence to the end of the Queue.
func Enqueue(as ...any) {
	Queue = append(Queue, as...)
}
