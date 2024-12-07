package wedge

import "github.com/stvmln86/wedge/wedge/atoms/word"

// Oper is a function that manipulates a Wedge.
type Oper func(*Wedge) error

// Opers is a map of all existing Oper functions.
var Opers = map[word.Word]Oper{
	"+": Add2,
}

// Add2 adds the top two items on the Wedge's Stack.
func Add2(w *Wedge) error {
	cs, err := w.Stack.PopN(2)
	if err != nil {
		return err
	}

	w.Stack.Push(cs[0] + cs[1])
	return nil
}
