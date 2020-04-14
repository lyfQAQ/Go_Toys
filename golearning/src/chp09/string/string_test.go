package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// string 是数据类型，只读的byte slic， len 为byte数
func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s)) // byte 数
	// s[1] = '3'  false	// 不可改变
	s = "\xE4\xB8\xA5" // 可存放任何二进制数据
	t.Log(s)

	c := []rune(s) //返回s的unicode 编码切片
	t.Log(c)
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",") // 分割字符串
	t.Log(s)
	t.Log(parts)
	t.Log(len(parts))
	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))
	for _, part := range parts {
		t.Log(part)
	}
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10) // 将数字10 转化为 stinrg 10
	t.Log("str" + s)

	val, _ := strconv.Atoi("10")
	t.Logf("%T %d", val, val)
}
