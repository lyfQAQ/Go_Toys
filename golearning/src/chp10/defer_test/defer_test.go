package defer_test

import (
	"fmt"
	"testing"
)

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear() // 延迟执行
	fmt.Println("Start")

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // 4 3 2 1 0
	}
}
