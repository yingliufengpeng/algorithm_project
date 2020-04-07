package datastruct

import (
	"errors"
	"fmt"
	"sync"
)

// QueueElement is an element of a linked queue.
type QueueElement struct {
	// Next and previous pointers in the doubly-linked queue of elements.
	next, prev *QueueElement
	// The queue to which this element belongs.
	queue *Queue
	// The value stored with this element.
	Value interface{}
}

// Queue represents a doubly linked queue.
type Queue struct {
	mutex *sync.RWMutex // mutex
	head  *QueueElement // head node
	tail  *QueueElement // tail node
	len   int           // current queue length excluding (this) sentinel element
	cap   int           // current queue cap
}

// NewQueue new a queue
func NewQueue(cap ...int) *Queue {
	_cap := 10
	if 0 != len(cap) {
		_cap = cap[0]
	}
	return &Queue{
		mutex: new(sync.RWMutex),
		head:  nil,
		tail:  nil,
		cap:   _cap,
	}
}

// Len return Queue size
func (q *Queue) Len() int {
	return q.len
}

// Cap return Queue cap
func (q *Queue) Cap() int {
	return q.cap
}

// Empty check Queue empty or not
func (q *Queue) Empty() bool {
	return q.len == 0
}

// Full check Queue Full or not
func (q *Queue) Full() bool {
	return q.len == q.cap
}

// Back get a element from queue tail
func (q *Queue) Back() *QueueElement {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.len > 0 {
		return q.tail
	}
	return nil
}

// Front get a element from queue head
func (q *Queue) Front() *QueueElement {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.len > 0 {
		return q.head
	}
	return nil
}

// Push push a element into queue tail
func (q *Queue) Push(v interface{}) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Full() {
		return errors.New("full queue")
	}
	if 0 == q.len {
		q.head = &QueueElement{
			next:  nil,
			prev:  nil,
			queue: q,
			Value: v,
		}
		q.tail = q.head
	} else {
		q.tail.next = &QueueElement{
			next:  nil,
			prev:  q.tail,
			queue: q,
			Value: v,
		}
		q.tail = q.tail.next
	}
	q.len++
	return nil
}

// Pop push a element into queue head
func (q *Queue) Pop() (*QueueElement, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Empty() {
		return nil, errors.New("empty queue")
	}
	e := q.head
	q.head = q.head.next
	q.head.prev = nil
	q.len--
	return e, nil
}

// String show queue
func (q *Queue) String() {
	if 0 == q.len {
		return
	}
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	ptr := q.head
	for ptr != nil && ptr.next != q.head {
		fmt.Printf("%v ", ptr.Value)
		ptr = ptr.next
	}
	fmt.Println()
}
