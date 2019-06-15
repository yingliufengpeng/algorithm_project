package stack

import (
	"fmt"
	"testing"
)




func TestPeng(t *testing.T) {

	var s = New()
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
