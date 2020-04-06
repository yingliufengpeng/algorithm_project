package datastruct

import (
	"errors"
	"fmt"
	"sync"
)

// ListElement is an element of a linked list.
type ListElement struct {
	// Next and previous pointers in the doubly-linked list of elements.
	next, prev *ListElement
	// The list to which this element belongs.
	list *List
	// The value stored with this element.
	Value interface{}
}

// List represents a doubly linked list.
type List struct {
	mutex *sync.RWMutex // mutex
	head  *ListElement  // head node
	tail  *ListElement  // tail node
	len   int           // current list length excluding (this) sentinel element
}

// NewList new a list
func NewList() *List {
	return &List{
		mutex: new(sync.RWMutex),
		head:  nil,
		tail:  nil,
	}
}

// Len return List size
func (l *List) Len() int {
	return l.len
}

// Back get a element from list tail
func (l *List) Back() *ListElement {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if l.len > 0 {
		return l.tail
	}
	return nil
}

// Front get a element from list head
func (l *List) Front() *ListElement {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if l.len > 0 {
		return l.head
	}
	return nil
}

// PushBack push a element into list tail
func (l *List) PushBack(v interface{}) interface{} {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if 0 == l.len {
		l.head = &ListElement{
			next:  nil,
			prev:  nil,
			list:  l,
			Value: v,
		}
		l.tail = l.head
	} else {
		l.tail.next = &ListElement{
			next:  nil,
			prev:  l.tail,
			list:  l,
			Value: v,
		}
		l.tail = l.tail.next
	}
	l.len++
	return v
}

// PushFront push a element into list head
func (l *List) PushFront(v interface{}) interface{} {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if 0 == l.len {
		l.tail = &ListElement{
			next:  nil,
			prev:  nil,
			list:  l,
			Value: v,
		}
		l.head = l.tail
	} else {
		l.head.prev = &ListElement{
			next:  l.head,
			prev:  nil,
			list:  l,
			Value: v,
		}
		l.head = l.head.prev
	}
	l.len++
	return v
}

// InsertBefore push a element into list before mark
func (l *List) InsertBefore(v interface{}, mark *ListElement) (e *ListElement) {
	if mark.list != l {
		return nil
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	e = &ListElement{
		next:  mark,
		prev:  mark.prev,
		list:  l,
		Value: v,
	}
	mark.prev = e
	e.prev.next = e
	l.len++
	return e
}

// InsertAfter push a element into list after a mark
func (l *List) InsertAfter(v interface{}, mark *ListElement) (e *ListElement) {
	if mark.list != l {
		return nil
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	e = &ListElement{
		next:  mark.next,
		prev:  mark,
		list:  l,
		Value: v,
	}
	mark.next = e
	e.next.prev = e
	l.len++
	return e
}

// Del delete a element from list
func (l *List) Del(v interface{}) (interface{}, error) {
	if 0 == l.len {
		return nil, errors.New("empty List")
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	ptr := l.head
	for ptr != nil {
		if v == ptr.Value {
			if nil == ptr.next {
				ptr.prev.next = nil
				l.tail = ptr.prev
			} else if nil == ptr.prev {
				l.head = ptr.next
				ptr.next.prev = nil
			} else {
				ptr.next.prev = ptr.prev
				ptr.prev.next = ptr.next
			}
			l.len--
			return v, nil
		}
		ptr = ptr.next
	}
	return nil, errors.New("find none")
}

// Find search a element from list
func (l *List) Find(v interface{}) (*ListElement, error) {
	if 0 == l.len {
		return nil, errors.New("empty List")
	}
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	ptr := l.head
	for ptr != nil {
		if v == ptr.Value {
			return ptr, nil
		}
		ptr = ptr.next
	}
	return nil, errors.New("find none")
}

// Concat merge List
func (l *List) Concat(ls ...*List) {
	for _, ll := range ls {
		if ll.Len() > 0 {
			if l.len == 0 {
				l.head = ll.head
				l.tail = ll.tail
				l.len = ll.Len()
			} else {
				ptr := ll.head
				for ptr != nil && ptr.next != ll.head {
					l.PushBack(ptr.Value)
					ptr = ptr.next
				}
			}
		}
	}
}

func (l *List) String() {
	if 0 == l.len {
		return
	}
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	ptr := l.head
	for ptr != nil && ptr.next != l.head {
		fmt.Printf("%v ", ptr.Value)
		ptr = ptr.next
	}
	fmt.Println()
}
