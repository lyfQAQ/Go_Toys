package operator_test

import "testing"

const (
	Readable   = 1 << iota // 001
	Writeable              // 010
	Executable             // 110
)

func TestBitClear(t *testing.T) {
	var a = 7          //0111
	a = a &^ Readable  // 去掉可读
	a = a &^ Writeable // 去掉可写
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) // false
	// t.Log(a == c)
	t.Log(a == d) // true
}
