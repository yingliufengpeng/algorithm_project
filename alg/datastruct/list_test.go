package datastruct

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()
	l.PushBack(1) // 1
	l.String()
	l.PushBack(2) // 1 2
	l.String()
	l.PushBack(3) // 1 2 3
	l.String()
	l.Del(1) // 2 3
	// l.Del(2) // 1 3
	// l.Del(3) // 1 2
	l.String()
	l.PushFront(0) // 0 1 2
	l.String()
	l.PushFront(-1) // -1 0 1 2
	l.String()
	mark, _ := l.Find(2)
	l.InsertBefore(999, mark)
	l.String()
	l.InsertAfter(666, mark)
	l.String()
	fmt.Println(l.Back().Value)
	fmt.Println(l.Front().Value)
	fmt.Println(l.Len())

	l2 := NewList()
	l2.PushBack(99)
	l2.PushBack(100)
	l2.PushFront(98)
	l2.String()

	l.Concat(l2)
	l2.String()
	l.String()

	l3 := NewList()
	l3.Concat(l, l2)
	l3.String()
}
