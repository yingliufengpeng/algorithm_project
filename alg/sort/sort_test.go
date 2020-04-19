package sort

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	d := []int{1, 2, 132, 4, 1, 0, 34}
	s := make([]interface{}, len(d))
	for i, v := range d {
		s[i] = v
	}
	st := NewSorter(0, func(a, b interface{}) bool {
		if a.(int) < b.(int) {
			return true
		}
		return false

	})
	st.SelectSort(&s)
	for i, v := range d {
		s[i] = v
	}
	st.InsertSort(&s)
	for i, v := range d {
		s[i] = v
	}
	st.QuickSort(&s)
	for i, v := range d {
		s[i] = v
	}
	st.MergeSort(&s)
}
