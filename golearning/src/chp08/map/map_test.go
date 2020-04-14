package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op }

	t.Log(m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] {
		t.Log("1 is existing")
	} else {
		t.Log("1 is not existing")
	}
	// 删除元素
	delete(mySet, 1)
}

func FindNoRepeatStringMaxLen(s string) (maxLen int) {
	// 存放字符最后一次出现的位置
	lastOccur := make(map[rune]int, 20)
	start := 0
	maxLen = 0
	for i, ch := range []rune(s) {
		if idx, ok := lastOccur[ch]; ok && idx >= start {
			start = idx + 1
		}
		if curLen := i - start + 1; curLen > maxLen {
			maxLen = curLen
		}
		lastOccur[ch] = i
	}
	return
}
func TestFindNoRepeatStringMaxLen(t *testing.T) {
	s := "中国人人啊"
	t.Log(FindNoRepeatStringMaxLen(s))
}
