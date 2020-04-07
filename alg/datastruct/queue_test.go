package datastruct

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue(5)
	fmt.Println(q.Empty())
	q.Push(0)
	q.Push(1)
	q.Push(3)
	q.String()
	q.Pop()
	q.String()
	fmt.Println(q.Front().Value)
	fmt.Println(q.Back().Value)
	q.Push(4)
	q.Push(5)
	q.String()
	q.Push(6)
	q.String()
	q.Push(7)
	q.Push(8)
	q.String()
}
