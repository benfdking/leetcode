package _go

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if minStack.GetMin() != -3 {
		t.Error()
	} // return -3
	minStack.Pop()
	if minStack.Top() != 0 {
		t.Error()
	} // return 0
	if minStack.GetMin() != -2 {
		t.Error()
	} // return -2
}

func BenchmarkConstructor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		minStack.GetMin()
		minStack.Pop()
		minStack.Top()
		minStack.GetMin()
	}
}

// 94.77 ns/op
