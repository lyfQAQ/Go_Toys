package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// b = a      Error: can not implicit conversion
	b = int64(a) // OK
	t.Log(a, b)

	var c int64 = 1
	var x MyInt // 不支持别名类型转换
	// x = c      Error: can not implicit conversion
	x = MyInt(c) // OK
	t.Log(x)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr := aPtr + 1     Error
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string // string为值类型
	t.Log("*" + s + "*")
	t.Log(len(s))
}
