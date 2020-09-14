package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generaterSlice(5)

	fmt.Println(slice)

	buildMaxHeap(slice)

	fmt.Println(slice)

}

func generaterSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(10)
	}
	return slice
}

func maxHeapify(nums []int, i int) {
	// left, right  都为最大堆
	left := i*2
	right := left+1
	var lagest int

	if len(nums) > left && nums[i] < nums[left] {
		lagest = left
	} else {
		lagest = i
	}

	if len(nums) > right && nums[i] < nums[right] {
		lagest = right
	}

	if lagest != i {
		nums[i], nums[lagest] = nums[lagest], nums[i]
		maxHeapify(nums, lagest)
	}
}

func buildMaxHeap(nums []int) {
	for i:= (len(nums))/2; i > 0; i-- {
		maxHeapify(nums, i);
	}
}


