package datastruct

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	bt := NewBitMap(100)
	bt.Set(10)
	fmt.Println(bt.Get(10))
	bt.Del(10)
	fmt.Println(bt.Get(10))
	fmt.Println(bt.Get(11))
	bt.Set(120)
	fmt.Println(bt.Set(1200000000000000000))
}
