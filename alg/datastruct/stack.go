package datastruct

import (
	"sync"
)

// StackElement is an element of a Stack.
type StackElement struct {
	// Next and previous pointers in the Stack of elements.
	prev *StackElement
	// The Stack to which this element belongs.
	stack *Stack
	// The value stored with this element.
	Value interface{}
}

// Stack represents a Stack.
type Stack struct {
	mutex *sync.RWMutex // mutex
	top   *StackElement // top node
	len   int           // current Stack length excluding (this) sentinel element
}

// NewStack Create a new stack
func NewStack() *Stack {
	return &Stack{
		mutex: new(sync.RWMutex),
		top:   nil,
	}
}

// Len return the number of items in the stack
func (s *Stack) Len() int {
	return s.len
}

// Peek view the top item on the stack
func (s *Stack) Peek() *StackElement {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.len == 0 {
		return nil
	}
	return s.top
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() (t *StackElement) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.len == 0 {
		return nil
	}
	t = s.top
	s.top = s.top.prev
	s.len--
	return t
}

// Push some value onto the top of the stack
func (s *Stack) Push(vv ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, v := range vv {
		n := &StackElement{
			s.top,
			s,
			v,
		}
		s.top = n
		s.len++
	}
}
