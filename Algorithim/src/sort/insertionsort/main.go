package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generaterSlice(10)

	fmt.Println(slice)

	insertionSrot(slice)

	fmt.Println(slice)

}

func generaterSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func insertionSrot(nums []int) {
	length := len(nums)
	for j := 1; j < length; j++ {
		key := nums[j]
		i := j - 1
		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			i--
		}
		nums[i+1] = key
	}
}
