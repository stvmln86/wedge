///////////////////////////////////////////////////////////////////////////////////////
//                            wedge · unit tests & helpers                           //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////////////
//                           part 0 · unit testing helpers                           //
///////////////////////////////////////////////////////////////////////////////////////

func assertOper(t *testing.T, s string, is ...int) {
	mockData(s)
	for len(Queue) > 0 {
		Evaluate(Dequeue())
	}

	if len(is) == 0 {
		assert.Empty(t, Stack)
	} else {
		assert.Equal(t, is, Stack)
	}
}

func mockData(s string, is ...int) {
	Queue = Parse(s)
	Stack = is
}

func mockStream(s string) *bytes.Buffer {
	Stdin = bytes.NewBufferString(s)
	Stdout = bytes.NewBuffer(nil)
	return Stdout.(*bytes.Buffer)
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part 1 · globals & constants                           //
///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////
//                              part 2 · input & output                              //
///////////////////////////////////////////////////////////////////////////////////////

func TestRead(t *testing.T) {
	// setup
	mockStream("test\n")

	// success
	i := Read()
	assert.Equal(t, int('t'), i)
}

func TestWrite(t *testing.T) {
	// setup
	b := mockStream("")

	// success
	Write(byte('t'))
	Write(int('t'))
	assert.Equal(t, "tt", b.String())
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part 3 · the stack & functions                          //
///////////////////////////////////////////////////////////////////////////////////////

func TestPeek(t *testing.T) {
	// setup
	mockData("", 1)

	// success
	i := Peek()
	assert.Equal(t, 1, i)
}

func TestPop(t *testing.T) {
	// setup
	mockData("", 1)

	// success
	i := Pop()
	assert.Equal(t, 1, i)
	assert.Empty(t, Stack)
}

func TestPush(t *testing.T) {
	// setup
	mockData("")

	// success
	Push(1)
	assert.Equal(t, []int{1}, Stack)
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part 4 · the queue & functions                          //
///////////////////////////////////////////////////////////////////////////////////////

func TestDequeue(t *testing.T) {
	// setup
	mockData("a")

	// success
	a := Dequeue()
	assert.Equal(t, "a", a)
	assert.Empty(t, Queue)
}

func TestDequeueTo(t *testing.T) {
	// setup
	mockData("a end")

	// success
	as := DequeueTo("end")
	assert.Equal(t, []any{"a"}, as)
	assert.Empty(t, Queue)
}

func TestEnqueue(t *testing.T) {
	// setup
	mockData("")

	// success
	Enqueue([]any{"a"})
	assert.Equal(t, []any{"a"}, Queue)
}

func TestInsert(t *testing.T) {
	// setup
	mockData("b")

	// success
	Insert([]any{"a"})
	assert.Equal(t, []any{"a", "b"}, Queue)
}

///////////////////////////////////////////////////////////////////////////////////////
//                           part 5 · parsing & evaluating                           //
///////////////////////////////////////////////////////////////////////////////////////

func TestParse(t *testing.T) {
	// success
	as := Parse("\t1 A\n")
	assert.Equal(t, []any{1, "a"}, as)
}

func TestEvaluate(t *testing.T) {
	// setup
	mockData("")

	// success - int
	Evaluate(1)
	assert.Equal(t, []int{1}, Stack)

	// success - string
	Evaluate("&")
	assert.Equal(t, []int{1, 1}, Stack)
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part 6 · operator functions                            //
///////////////////////////////////////////////////////////////////////////////////////

func TestInitOpers(t *testing.T) {
	// setup
	b := mockStream("test\n")

	// success - mathematical operators
	assertOper(t, "1 2 +", 3)
	assertOper(t, "1 2 -", 1)
	assertOper(t, "2 3 *", 6)
	assertOper(t, "3 6 /", 2)
	assertOper(t, "2 5 %", 1)

	// success - stack operators
	assertOper(t, "1 &", 1, 1)
	assertOper(t, "1 #", 1, 1)
	assertOper(t, "1 2 ~", 2, 1)
	assertOper(t, "1 2 3 @", 2, 3, 1)

	// success - input/output operators
	assertOper(t, "116 .")
	assertOper(t, ",", int('t'))
	assert.Equal(t, "t", b.String())

	// logic operators
	assertOper(t, "0 {? 1 ?} · 1 {? 1 ?}", 1)
	assertOper(t, "0 {# 1 #} · 2 {# 1 #}", 1, 1)
	assertOper(t, "{= t 1 =} t", 1)

	// setup
	b.Reset()

	// system operators
	assertOper(t, "·")
	assertOper(t, "1 dump", 1)
	assert.Equal(t, ": [1]\n", b.String())

	assertOper(t, "exit")
	assert.False(t, Running)
	Running = true

	assertOper(t, "38 32 49 eval", 1, 1)
}

///////////////////////////////////////////////////////////////////////////////////////
//                            part 7 · top-level commands                            //
///////////////////////////////////////////////////////////////////////////////////////

func TestRunREPL(t *testing.T) {
	// setup
	mockData("")
	b := mockStream("1 & exit\n")

	// success
	RunREPL()
	assert.Equal(t, []int{1, 1}, Stack)
	assert.Equal(t, "> : [1 1]\n", b.String())
}

///////////////////////////////////////////////////////////////////////////////////////
//                              part 8 · main functions                              //
///////////////////////////////////////////////////////////////////////////////////////

func init() {
	InitOpers()
}
