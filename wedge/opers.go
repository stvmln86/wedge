package wedge

// Oper is a global operator function.
type Oper func()

// Opers is a global map of all existing Oper functions.
var Opers = make(map[string]Oper)

// init initialises the Opers map.
func init() {
	// Mathematical functions.
	Opers["+"] = func() { Push(Pop() + Pop()) }
	Opers["-"] = func() { Push(Pop() - Pop()) }
	Opers["*"] = func() { Push(Pop() * Pop()) }
	Opers["/"] = func() { Push(Pop() / Pop()) }
	Opers["%"] = func() { Push(Pop() % Pop()) }

	// Stack functions.
	Opers["&"] = func() { Push(Peek()) }
	Opers["#"] = func() { Push(len(Stack)) }
	Opers["~"] = func() { Push(Pop(), Pop()) }
	Opers["@"] = func() {
		a, b, c := Pop(), Pop(), Pop()
		Push(b, a, c)
	}

	// Input/output functions.
	Opers["."] = func() { Write(Pop()) }
	Opers[","] = func() { Push(Read()) }

	// Logic functions.
	Opers["?["] = func() {
		as := DequeueTo("]?")
		if Pop() != 0 {
			EvaluateSlice(as)
		}
	}
}
