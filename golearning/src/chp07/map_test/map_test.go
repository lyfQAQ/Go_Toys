package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 7}
	t.Log(m1[1])
	t.Logf("len m1 = %d", len(m1))

	// key不存在时 value初始化为0
	if _, ok := m1[3]; ok {
		t.Log("Key 3's value is existing")
	} else {
		t.Log("Key 3's value is not existing")
	}

	m2 := map[int]int{}
	m2[2] = 19
	t.Logf("len m2 = %d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3 = %d", len(m3))
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 7}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
