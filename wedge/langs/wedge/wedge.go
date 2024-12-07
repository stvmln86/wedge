// Package wedge implements the Wedge type and Oper functions.
package wedge

import (
	"fmt"
	"io"

	"github.com/stvmln86/wedge/wedge/atoms/cell"
	"github.com/stvmln86/wedge/wedge/atoms/word"
	"github.com/stvmln86/wedge/wedge/langs/parse"
	"github.com/stvmln86/wedge/wedge/langs/queue"
	"github.com/stvmln86/wedge/wedge/langs/stack"
)

// Wedge is a top-leve program container and controller.
type Wedge struct {
	Queue  *queue.Queue
	Stack  *stack.Stack
	Reader io.Reader
	Writer io.Writer
}

// New returns a new empty Wedge.
func New(r io.Reader, w io.Writer) *Wedge {
	return &Wedge{queue.New(), stack.New(), r, w}
}

// Evaluate evaluates the next Word in the Wedge's Queue.
func (w *Wedge) Evaluate() error {
	a, err := w.Queue.Dequeue()
	if err != nil {
		return err
	}

	switch a := a.(type) {
	case cell.Cell:
		w.Stack.Push(a)
		return nil
	case word.Word:
		f, b := Opers[a]
		if !b {
			return fmt.Errorf("invalid Word %q", a)
		}

		return f(w)
	default:
		return fmt.Errorf("invalid Atom %q", a)
	}
}

// EvaluateAll evaluates all Words in the Wedge's Queue.
func (w *Wedge) EvaluateAll() error {
	for !w.Queue.Empty() {
		if err := w.Evaluate(); err != nil {
			return err
		}
	}

	return nil
}

// EvaluateString parses, enqueues and evaluates a string.
func (w *Wedge) EvaluateString(s string) error {
	as, err := parse.ParseString(s)
	if err != nil {
		return err
	}

	w.Queue.EnqueueAll(as)
	return w.EvaluateAll()
}
