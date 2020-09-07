package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := generaterSlice(100000000)
	fmt.Println(nums)
	res := peekFindDivide(nums, 0, len(nums)-1)
	fmt.Println(res)
}

func peekFindDivide(nums []int, left int, right int) int {

	mid := left + (right-left)/2
	if (mid+1) < len(nums) && nums[mid] <= nums[mid+1] {
		return peekFindDivide(nums, mid+1, right)
	} else if (mid-1) >= 0 && nums[mid] <= nums[mid-1] {
		return peekFindDivide(nums, left, mid)
	} else {
		return nums[mid]
	}
}

func generaterSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(10000)
	}
	return slice
}
