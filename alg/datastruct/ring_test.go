package datastruct

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := NewRing(5)
	fmt.Println(r.Empty())
	r.Push(0)
	r.Push(1)
	r.Push(3)
	r.String()
	r.Pop()
	r.String()
	fmt.Println(r.Front().Value)
	fmt.Println(r.Back().Value)
	r.Push(4)
	r.Push(5)
	r.String()
	r.Push(6)
	r.String()
	r.Push(7)
	r.Push(8)
	r.String()
}
