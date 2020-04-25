package datastruct

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue(50)
	fmt.Println(q.Empty())
	q.Push(0, 10)
	q.Push(1, 5)
	q.Push(3, 11)
	q.Push(3, -1)
	q.Push(4, 11)
	q.Push(99, -2)
	q.String()
	fmt.Println(q.Peek())
	for {
		if d, ok := q.Pop(); nil == ok {
			fmt.Println(d)
		} else {
			break
		}
	}
}
