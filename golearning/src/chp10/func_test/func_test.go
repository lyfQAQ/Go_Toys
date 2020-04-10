package func_test

import "testing"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxFunc(t *testing.T) {
	t.Log(max(4, 5))
}

// 返回多个值
// go 中很多Package都会返回两个值，一个是正常值，一个是错误

func multiVal(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	var err bool
	var val int
	val, err = m[key]
	return val, err
}
func TestFuncMultiVal(t *testing.T) {
	v, e := multiVal("one")
	t.Log(v, e) // 1 true

	v, e = multiVal("four")
	t.Log(v, e) // 0 false

	// 通常的用法
	if v, e := multiVal("four"); e {
		t.Log("existing", v)
	} else {
		t.Log("not existing")
	}
}

// 可变长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(6, 6, 43))

	// 传切片
	arr := []int{1, 2, 3, 4, 5}
	t.Log(Sum(arr...))
}
