package datastruct

import (
	"errors"
	"fmt"
	"sync"
)

// PriorityQueueElement is an element of a linked queue.
type PriorityQueueElement struct {
	// The queue to which this element belongs.
	queue *PriorityQueue
	// The value stored with this element.
	Value    interface{}
	Priority int
}

// PriorityQueue represents a doubly linked queue.
// Using doubly linked
// insert: O(n) -> Push
// delete: O(1) -> Pop
type PriorityQueue struct {
	mutex *sync.RWMutex           // mutex
	data  []*PriorityQueueElement // head node
	len   int                     // current queue length excluding (this) sentinel element
	cap   int                     // current queue cap
}

// NewPriorityQueue new a queue
func NewPriorityQueue(cap ...int) *PriorityQueue {
	_cap := 10
	if 0 != len(cap) {
		_cap = cap[0]
	}
	return &PriorityQueue{
		mutex: new(sync.RWMutex),
		data:  make([]*PriorityQueueElement, _cap),
		cap:   _cap,
	}
}

// Full check PriorityQueue Full or not
func (q *PriorityQueue) parent(i int) (int, bool) {
	index := i >> 1
	if index > q.len || index < 0 {
		return -1, false
	}
	return index, true
}

// Full check PriorityQueue Full or not
func (q *PriorityQueue) lchild(i int) (int, bool) {
	index := (i << 1) + 1
	if index > q.len || index < 0 {
		return -1, false
	}
	return index, true
}

// Full check PriorityQueue Full or not
func (q *PriorityQueue) rchild(i int) (int, bool) {
	index := (i + 1) << 1
	if index > q.len || index < 0 {
		return -1, false
	}
	return index, true
}

// Len return PriorityQueue size
func (q *PriorityQueue) Len() int {
	return q.len
}

// Cap return PriorityQueue cap
func (q *PriorityQueue) Cap() int {
	return q.cap
}

// Empty check PriorityQueue empty or not
func (q *PriorityQueue) Empty() bool {
	return q.len == 0
}

// Full check PriorityQueue Full or not
func (q *PriorityQueue) Full() bool {
	return q.len == q.cap
}

// Peek get a element from queue
func (q *PriorityQueue) Peek() *PriorityQueueElement {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.len > 0 {
		return q.data[0]
	}
	return nil
}

// insert insert a element in a right position
func (q *PriorityQueue) insert(e *PriorityQueueElement, p int) {
	q.data[q.len] = e
	for i := q.len; i > 0; {
		index, _ := q.parent(i)
		if q.data[index].Priority >= q.data[i].Priority {
			break
		}
		q.data[index], q.data[i] = q.data[i], q.data[index]
		i = index
	}
	q.len++
}

// insert insert a element in a right position
func (q *PriorityQueue) delete() *PriorityQueueElement {
	e := q.data[0]
	q.data[0] = q.data[q.len-1]
	q.len--
	i := 0
	for 2*i <= q.len {
		index := 2 * i
		if index < q.len && q.data[index].Priority < q.data[index+1].Priority {
			index++
		}
		if q.data[i].Priority >= q.data[index].Priority {
			break
		}
		q.data[i], q.data[index] = q.data[index], q.data[i]
		i = index
	}
	return e
}

// Push push a element into queue
func (q *PriorityQueue) Push(v interface{}, p int) (*PriorityQueueElement, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Full() {
		return nil, errors.New("full queue")
	}
	e := &PriorityQueueElement{
		queue:    q,
		Value:    v,
		Priority: p,
	}
	q.insert(e, p)
	return e, nil
}

// Pop push a element into queue head
func (q *PriorityQueue) Pop() (*PriorityQueueElement, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.Empty() {
		return nil, errors.New("empty queue")
	}
	e := q.delete()
	return e, nil
}

// String show queue
func (q *PriorityQueue) String() {
	if 0 == q.len {
		return
	}
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	for i := 0; i < q.len; i++ {
		fmt.Println(q.data[i])
	}
	fmt.Println()
}
