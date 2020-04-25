package datastruct

import (
	"errors"
	"fmt"
	"sync"
)

// PriorityQueue2Element is an element of a linked queue.
type PriorityQueue2Element struct {
	// Next and previous pointers in the doubly-linked queue of elements.
	next, prev *PriorityQueue2Element
	// The queue to which this element belongs.
	queue *PriorityQueue2
	// The value stored with this element.
	Value    interface{}
	Priority int
}

// PriorityQueue2 represents a doubly linked queue.
// Using doubly linked
// insert: O(n) -> Push
// delete: O(1) -> Pop
type PriorityQueue2 struct {
	mutex *sync.RWMutex          // mutex
	head  *PriorityQueue2Element // head node
	tail  *PriorityQueue2Element // tail node
	len   int                    // current queue length excluding (this) sentinel element
	cap   int                    // current queue cap
}

// NewPriorityQueue2 new a queue
func NewPriorityQueue2(cap ...int) *PriorityQueue2 {
	_cap := 10
	if 0 != len(cap) {
		_cap = cap[0]
	}
	return &PriorityQueue2{
		mutex: new(sync.RWMutex),
		head:  nil,
		tail:  nil,
		cap:   _cap,
	}
}

// Len return PriorityQueue2 size
func (q *PriorityQueue2) Len() int {
	return q.len
}

// Cap return PriorityQueue2 cap
func (q *PriorityQueue2) Cap() int {
	return q.cap
}

// Empty check PriorityQueue2 empty or not
func (q *PriorityQueue2) Empty() bool {
	return q.len == 0
}

// Full check PriorityQueue2 Full or not
func (q *PriorityQueue2) Full() bool {
	return q.len == q.cap
}

// Peek get a element from queue
func (q *PriorityQueue2) Peek() *PriorityQueue2Element {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.len > 0 {
		return q.head
	}
	return nil
}

// insert insert a element in a right position
func (q *PriorityQueue2) insert(e *PriorityQueue2Element, p int) {
	if p >= q.head.Priority {
		e.next = q.head
		q.head.prev = e
		q.head = e
	} else if p <= q.tail.Priority {
		e.prev = q.tail
		q.tail.next = e
		q.tail = e
	} else {
		cur := q.head
		for cur.Priority > p {
			cur = cur.prev
		}
		e.next = cur
		e.prev = cur.prev
		cur.prev.next = e
		cur.prev = e
	}
}

// Push push a element into queue
func (q *PriorityQueue2) Push(v interface{}, p int) (*PriorityQueue2Element, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Full() {
		return nil, errors.New("full queue")
	}
	e := &PriorityQueue2Element{
		next:     nil,
		prev:     nil,
		queue:    q,
		Value:    v,
		Priority: p,
	}
	if 0 == q.len {
		q.head = e
		q.tail = q.head
		q.head.next = q.tail
		q.tail.prev = q.head
	} else {
		q.insert(e, p)
	}
	q.len++
	return e, nil
}

// Pop push a element into queue head
func (q *PriorityQueue2) Pop() (*PriorityQueue2Element, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Empty() {
		return nil, errors.New("empty queue")
	}
	e := q.head
	q.len--
	if 0 < q.len {
		q.head = q.head.next
		q.head.prev = nil
	}
	return e, nil
}

// String show queue
func (q *PriorityQueue2) String() {
	if 0 == q.len {
		return
	}
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	ptr := q.head
	for ptr != nil && ptr.next != q.head {
		fmt.Println(ptr.Value, ptr.Priority)
		ptr = ptr.next
	}
	fmt.Println()
}
