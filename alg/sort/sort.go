package sort

import "errors"

// Method
const (
	Select = iota
	Insert
	Quick
	Merge
)

// Sorter sorter
type Sorter struct {
	Method byte
	Data   *[]interface{}
	cmp    func(a, b interface{}) bool
}

// NewSorter new a sorter
func NewSorter(md byte, f func(a, b interface{}) bool) *Sorter {
	return &Sorter{
		Method: md,
		Data:   nil,
		cmp:    f,
	}
}

// SelectSort select sort
func (st *Sorter) SelectSort(data *[]interface{}) error {
	if nil == data && nil == st.Data {
		return errors.New("sorter have no data")
	}
	dlen := len(*data)
	for i := 0; i < dlen; i++ {
		letter := i
		for j := i + 1; j < dlen; j++ {
			if st.cmp((*data)[j], (*data)[letter]) {
				letter = j
			}
		}
		(*data)[i], (*data)[letter] = (*data)[letter], (*data)[i]
	}
	return nil
}

// InsertSort insert sort
func (st *Sorter) InsertSort(data *[]interface{}) error {
	if nil == data && nil == st.Data {
		return errors.New("sorter have no data")
	}
	dlen := len(*data)
	for i := 1; i < dlen; i++ {
		for j := i; j > 0 && st.cmp((*data)[j], (*data)[j-1]); j-- {
			(*data)[j], (*data)[j-1] = (*data)[j-1], (*data)[j]
		}
	}
	return nil
}

// QuickSort quick sort
func (st *Sorter) QuickSort(data *[]interface{}) error {
	if nil == data && nil == st.Data {
		return errors.New("sorter have no data")
	}
	st.quickSort(data, 0, len(*data)-1)
	return nil
}

func (st *Sorter) quickSort(data *[]interface{}, left int, right int) {
	if left < right {
		index := st.getIndex(data, left, right)
		st.quickSort(data, 0, index-1)
		st.quickSort(data, index+1, right)
	}
}

func (st *Sorter) getIndex(data *[]interface{}, left int, right int) int {
	key := (*data)[left]
	for left < right {
		for left < right && !st.cmp((*data)[right], key) {
			right--
		}
		(*data)[left] = (*data)[right]
		for left < right && st.cmp((*data)[left], key) {
			left++
		}
		(*data)[right] = (*data)[left]
	}
	(*data)[left] = key
	return left
}

// MergeSort merge sort
func (st *Sorter) MergeSort(data *[]interface{}) error {
	if nil == data && nil == st.Data {
		return errors.New("sorter have no data")
	}
	st.mergeSort(data, 0, len(*data)-1)
	return nil
}

func (st *Sorter) mergeSort(data *[]interface{}, left int, right int) {
	if left < right {
		mid := left + (right-left)/2
		st.mergeSort(data, 0, mid)
		st.mergeSort(data, mid+1, right)
		st.merge(data, left, mid, right)
	}
}

func (st *Sorter) merge(data *[]interface{}, left int, mid int, right int) {
	data2 := make([]interface{}, right-left+1)
	i := left    // 第1个有序区的索引
	j := mid + 1 // 第2个有序区的索引
	k := 0       // 临时区域的索引
	for i <= mid && j <= right {
		if st.cmp((*data)[i], (*data)[j]) {
			data2[k] = (*data)[i]
			k++
			i++
		} else {
			data2[k] = (*data)[j]
			k++
			j++
		}
	}
	for i <= mid {
		data2[k] = (*data)[i]
		k++
		i++
	}
	for j <= right {
		data2[k] = (*data)[j]
		k++
		j++
	}
	for i := 0; i < k; i++ {
		(*data)[left+i] = data2[i]
	}
}
