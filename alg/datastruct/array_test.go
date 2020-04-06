package datastruct

import (
	"reflect"
	"testing"
)

// Equal checks if two input are deeply equal.
func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}

func TestArray(t *testing.T) {
	arr := NewArray(10)

	Equal(t, 0, arr.Len())

	arr.Push(1)
	arr.Push(2, 3, 4)
	Equal(t, 4, arr.Len())

	arr.Pop()
	Equal(t, 3, arr.Len())
	arr.Pop(2)
	Equal(t, 1, arr.Len())
	arr.Pop(100)
	Equal(t, 0, arr.Len())

	arr.Unshift(0)
	Equal(t, 1, arr.Len())
	arr.Unshift(-4, -3, -2, -1)
	Equal(t, 5, arr.Len())

	arr.Shift()
	Equal(t, 4, arr.Len())
	arr.Shift(2)
	Equal(t, 2, arr.Len())
	arr.Shift(100)
	Equal(t, 0, arr.Len())
	arr.Push(0)

	arr.Get(0)
	arr.Push(1)
	arr.Push(2)
	arr.Set(1, 3)
	arr.Push(2)

	Equal(t, nil, arr.Find(1))
	Equal(t, 3, arr.Find(3))
	arr.FindAll(2)
	arr.Del(1)

	arr2 := NewArray(10)
	arr2.Push(90, 91)

	arr.Concat(arr2)

	arr3 := NewArray(10)
	arr3.Push(92, 93)
	arr4 := NewArray(10)
	arr4.Push(94, 95, 96)

	arr.Concat(arr3, arr4)

	arr.Concat()
	Equal(t, arr, arr)

	arr.Reverse()
}
