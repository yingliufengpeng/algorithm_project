package stack

import (
	"testing"
	"fmt"
)

var s = New()

func TestStringPush(t *testing.T) {
	s.Push("H")
	fmt.Println(s.Peek())
	s.Push("e")
	fmt.Println(s.Peek())
	s.Push("l")
	fmt.Println(s.Peek())
	s.Push("l")
	fmt.Println(s.Peek())
	s.Push("o")
	fmt.Println(s.Peek())
	fmt.Println("=======TestStringPush=======")
}

func TestStringPop(t *testing.T) {
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	fmt.Println("=======TestStringPop=======")
}

func TestStuctPush(t *testing.T) {
	type d struct {
		name string
		age  int
	}
	s.Push(&d{name:"a", age:10})
	fmt.Println(s.Peek())
	s.Push(&d{name:"b", age:11})
	fmt.Println(s.Peek())
	s.Push(&d{name:"c", age:12})
	fmt.Println(s.Peek())
	s.Push(&d{name:"d", age:13})
	fmt.Println(s.Peek())
	s.Push(&d{name:"e", age:14})
	fmt.Println(s.Peek())
	fmt.Println("=======TestStuctPush=======")
}

func TestStuctPop(t *testing.T) {
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	fmt.Println("=======TestStuctPop=======")
}

