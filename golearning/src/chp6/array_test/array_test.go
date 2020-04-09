package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2]) // 0 0
	t.Log(arr1, arr3)
}

// 遍历数组
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	for _, e := range arr3 {
		t.Log(e)
	}
}

// 数组截取
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	// arr3_sec := arr3[:3]  //[1,2,3]
	arr3_sec := arr3[3:] // [4,5]
	t.Log(arr3_sec)
}
