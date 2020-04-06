package datastruct

import (
	"reflect"
)

// ArrayElement support all Type
type ArrayElement interface{}

// Array array Type
type Array struct {
	_data []ArrayElement
	_len  int
}

// NewArray new a array
func NewArray(size ...int) *Array {
	_size := 10
	if 0 != len(size) {
		_size = size[0]
	}
	return &Array{_data: make([]ArrayElement, 0, _size)}
}

// Len Array use size
func (arr *Array) Len() int {
	return arr._len
}

// Push insert <d...> element into Array tail
func (arr *Array) Push(e ...ArrayElement) {
	arr._data = append(arr._data, e...)
	arr._len += len(e)
}

// Pop del <size> element from Array tail
func (arr *Array) Pop(size ...int) (a *Array) {
	_size := 1
	if 0 != len(size) {
		_size = size[0]
		if _size > arr._len {
			_size = arr._len
		}
	}
	a = NewArray()
	a._data = arr._data[len(arr._data)-_size:]
	a._len = _size
	arr._data = arr._data[:len(arr._data)-_size]
	arr._len -= _size
	return a
}

// Unshift insert <d...> element into Array head
func (arr *Array) Unshift(e ...ArrayElement) {
	arr._data = append(append([]ArrayElement{}, e...), arr._data...)
	arr._len += len(e)
}

// Shift del <size> element input Array head
func (arr *Array) Shift(size ...int) (a *Array) {
	_size := 1
	if 0 != len(size) {
		_size = size[0]
		if _size > arr._len {
			_size = arr._len
		}
	}
	a = NewArray()
	a._data = arr._data[0:_size]
	a._len = _size
	arr._data = arr._data[_size:]
	arr._len -= _size
	return a
}

// Set modify element in Array a postion
func (arr *Array) Set(pos int, e ArrayElement) {
	if pos > arr._len-1 || pos < 0 {
		return
	}
	arr._data[pos] = e
}

// Del delete element in Array a postion
func (arr *Array) Del(pos int) (e ArrayElement) {
	if pos > arr._len-1 || pos < 0 {
		return nil
	}
	e = arr._data[pos]
	arr._data = append(arr._data[0:pos], arr._data[pos+1:]...)
	arr._len--
	return e
}

// Get get element in Array a postion
func (arr *Array) Get(pos int) ArrayElement {
	if pos > arr._len-1 || pos < 0 {
		return nil
	}
	return arr._data[pos]
}

// Find find element in Array, only find first match
func (arr *Array) Find(e ArrayElement) ArrayElement {
	for _, _d := range arr._data {
		if reflect.DeepEqual(_d, e) {
			return _d
		}
	}
	return nil
}

// FindAll find element in Array, will return all match
func (arr *Array) FindAll(d ArrayElement) (e *Array) {
	e = NewArray()
	for _, _d := range arr._data {
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
	return arr.Get(arr._len - 1)
}

// Concat merge Array
func (arr *Array) Concat(arrs ...*Array) {
	for _, ar := range arrs {
		if ar.Len() > 0 {
			arr._data = append(arr._data, ar._data...)
			arr._len += ar.Len()
		}
	}
}

// Reverse reverse Array
func (arr *Array) Reverse() {
	for i := range arr._data {
		if i < arr._len/2 {
			arr._data[i], arr._data[arr._len-1-i] = arr._data[arr._len-1-i], arr._data[i]
		} else {
			break
		}
	}
}
