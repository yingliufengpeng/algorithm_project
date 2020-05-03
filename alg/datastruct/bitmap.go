package datastruct

import "errors"

// BitMapCAP bit map base capacity
var BitMapCAP int = 512

// BitMap bit map
type BitMap struct {
	// bit data
	data []byte
	// bit map max capacity
	cap int
}

// NewBitMap new a bit map
func NewBitMap(size ...int) *BitMap {
	_size := BitMapCAP
	if 0 != len(size) {
		_size = size[0]
	}
	return &BitMap{
		data: make([]byte, _size),
		cap:  _size,
	}
}

// Set set a number into bit map
func (bt *BitMap) Set(n int) error {
	index := n >> 3
	if index > bt.cap {
		return errors.New("over flow capacity")
	}
	par := n & 0x07
	bt.data[index] |= (1 << par)
	return nil
}

// Get set a number into bit map
func (bt *BitMap) Get(n int) bool {
	index := n >> 3
	if index > bt.cap {
		return false
	}
	par := n & 0x07
	return bt.data[index]&(1<<par) != 0
}

// Del clear a number from bit map
func (bt *BitMap) Del(n int) {
	index := n >> 3
	par := n & 0x07
	bt.data[index] &= ^(1 << par)
}

func (bt *BitMap) String() string {
	return ""
}
