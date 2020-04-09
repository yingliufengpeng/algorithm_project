package datastruct

import (
	"fmt"
	"testing"
)

func TestTuopu(t *testing.T) {
	tp := NewTuopu()
	tp.Push("A", "B", "C")
	fmt.Println(tp.Get("A"))
	fmt.Println(tp.Get("B"))
	tp.Push("B", "C", "D")
	tp.Push("A", "E", "F")
	tp.Push("C", "E", "D")
	tp.Push("F", "A", "G")
	fmt.Println(tp.node)
	fmt.Println(tp.link)
	fmt.Println(tp.Get("A"))
}
