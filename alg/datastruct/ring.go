package datastruct

import (
	"errors"
	"fmt"
	"sync"
)

// RingElement is an element of a linked ring.
type RingElement struct {
	// Next and previous pointers in the doubly-linked ring of elements.
	next, prev *RingElement
	// The ring to which this element belongs.
	ring *Ring
	// The value stored with this element.
	Value interface{}
}

// Ring represents a doubly linked ring.
type Ring struct {
	mutex *sync.RWMutex // mutex
	head  *RingElement  // head node
	tail  *RingElement  // tail node
	len   int           // current ring length excluding (this) sentinel element
	cap   int           // current ring cap
}

// NewRing new a ring
func NewRing(cap ...int) *Ring {
	_cap := 10
	if 0 != len(cap) {
		_cap = cap[0]
	}
	return &Ring{
		mutex: new(sync.RWMutex),
		head:  nil,
		tail:  nil,
		cap:   _cap,
	}
}

// Len return Ring size
func (r *Ring) Len() int {
	return r.len
}

// Cap return Ring cap
func (r *Ring) Cap() int {
	return r.cap
}

// Empty check Ring empty or not
func (r *Ring) Empty() bool {
	return r.len == 0
}

// Full check Ring Full or not
func (r *Ring) Full() bool {
	return r.len == r.cap
}

// Back get a element from ring tail
func (r *Ring) Back() *RingElement {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.len > 0 {
		return r.tail
	}
	return nil
}

// Front get a element from ring head
func (r *Ring) Front() *RingElement {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.len > 0 {
		return r.head
	}
	return nil
}

// Push push a element into ring tail
func (r *Ring) Push(v interface{}) (*RingElement, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.Full() {
		return nil, errors.New("full ring")
	}
	var e *RingElement
	if 0 == r.len {
		e = &RingElement{
			next:  nil,
			prev:  nil,
			ring:  r,
			Value: v,
		}
		r.head = e
		r.tail = r.head
		r.head.next = r.tail
		r.tail.prev = r.head
	} else {
		e = &RingElement{
			next:  nil,
			prev:  r.tail,
			ring:  r,
			Value: v,
		}
		r.tail.next = e
		r.tail = r.tail.next
		r.tail.next = r.head
		r.head.prev = r.tail
	}
	r.len++
	return e, nil
}

// Pop push a element into ring head
func (r *Ring) Pop() (*RingElement, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.Empty() {
		return nil, errors.New("empty ring")
	}
	e := r.head
	r.head = r.head.next
	r.tail.next = r.head
	r.head.prev = r.tail
	r.len--
	return e, nil
}

// String show ring
func (r *Ring) String() {
	if 0 == r.len {
		return
	}
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	ptr := r.head
	for ptr != nil && ptr.next != r.head {
		fmt.Printf("%v ", ptr.Value)
		ptr = ptr.next
	}
	fmt.Println()
}
