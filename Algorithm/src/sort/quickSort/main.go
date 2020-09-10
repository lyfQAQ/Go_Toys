package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generaterSlice(1)

	fmt.Println(slice)

	res := quickSort(slice, 0, 0)

	fmt.Println(res)
}

func generaterSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func quickSort(nums []int, begin, end int) []int {
	if begin < end {
		i := partital(nums, begin, end)
		quickSort(nums, begin, i-1)
		quickSort(nums, i+1, end)
	}
	return nums
}

// [begin, end]
func partital(nums []int, begin, end int) int {
	key := nums[begin]
	i := begin
	// i 始终表示最后一个<=key的数的下标
	for j := begin + 1; j <= end; j++ {
		if nums[j] <= key {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 最后一个<=key的数和key交换
	nums[begin], nums[i] = nums[i], nums[begin]
	return i
}
