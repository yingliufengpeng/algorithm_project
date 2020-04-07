package datastruct

import (
	"reflect"
)

// ArrayElement support all Type
type ArrayElement interface{}

// Array array Type
type Array struct {
	value []ArrayElement
	len   int
}

// NewArray new a array
func NewArray(size ...int) *Array {
	_size := 10
	if 0 != len(size) {
		_size = size[0]
	}
	return &Array{value: make([]ArrayElement, 0, _size)}
}

// Len Array use size
func (arr *Array) Len() int {
	return arr.len
}

// Push insert <d...> element into Array tail
func (arr *Array) Push(e ...ArrayElement) {
	arr.value = append(arr.value, e...)
	arr.len += len(e)
}

// Pop del <size> element from Array tail
func (arr *Array) Pop(size ...int) (a *Array) {
	_size := 1
	if 0 != len(size) {
		_size = size[0]
		if _size > arr.len {
			_size = arr.len
		}
	}
	a = NewArray()
	a.value = arr.value[len(arr.value)-_size:]
	a.len = _size
	arr.value = arr.value[:len(arr.value)-_size]
	arr.len -= _size
	return a
}

// Unshift insert <d...> element into Array head
func (arr *Array) Unshift(e ...ArrayElement) {
	arr.value = append(append([]ArrayElement{}, e...), arr.value...)
	arr.len += len(e)
}

// Shift del <size> element input Array head
func (arr *Array) Shift(size ...int) (a *Array) {
	_size := 1
	if 0 != len(size) {
		_size = size[0]
		if _size > arr.len {
			_size = arr.len
		}
	}
	a = NewArray()
	a.value = arr.value[0:_size]
	a.len = _size
	arr.value = arr.value[_size:]
	arr.len -= _size
	return a
}

// Set modify element in Array a postion
func (arr *Array) Set(pos int, e ArrayElement) {
	if pos > arr.len-1 || pos < 0 {
		return
	}
	arr.value[pos] = e
}

// Del delete element in Array a postion
func (arr *Array) Del(pos int) (e ArrayElement) {
	if pos > arr.len-1 || pos < 0 {
		return nil
	}
	e = arr.value[pos]
	arr.value = append(arr.value[0:pos], arr.value[pos+1:]...)
	arr.len--
	return e
}

// Get get element in Array a postion
func (arr *Array) Get(pos int) ArrayElement {
	if pos > arr.len-1 || pos < 0 {
		return nil
	}
	return arr.value[pos]
}

// Find find element in Array, only find first match
func (arr *Array) Find(e ArrayElement) ArrayElement {
	for _, _d := range arr.value {
		if reflect.DeepEqual(_d, e) {
			return _d
		}
	}
	return nil
}

// FindAll find element in Array, will return all match
func (arr *Array) FindAll(d ArrayElement) (e *Array) {
	e = NewArray()
	for _, _d := range arr.value {
		if reflect.DeepEqual(_d, d) {
			e.Push(_d)
		}
	}
	return e
}

// Head get element in Array head
func (arr *Array) Head() ArrayElement {
	return arr.Get(0)
}

// Tail get element in Array tail
func (arr *Array) Tail() ArrayElement {
	return arr.Get(arr.len - 1)
}

// Concat merge Array
func (arr *Array) Concat(arrs ...*Array) {
	for _, ar := range arrs {
		if ar.Len() > 0 {
			arr.value = append(arr.value, ar.value...)
			arr.len += ar.Len()
		}
	}
}

// Reverse reverse Array
func (arr *Array) Reverse() {
	for i := range arr.value {
		if i < arr.len/2 {
			arr.value[i], arr.value[arr.len-1-i] = arr.value[arr.len-1-i], arr.value[i]
		} else {
			break
		}
	}
}
