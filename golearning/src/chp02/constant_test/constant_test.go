package constant_test

import "testing"

const (
	Monday    = iota + 1 // 1
	Tuesday              // 2
	Wednesday            // 3
)

const (
	Readable   = 1 << iota // 001
	Writeable              // 010
	Executable             // 110
)

func TestConstantTry1(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry2(t *testing.T) {
	var a = 7 //111
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
