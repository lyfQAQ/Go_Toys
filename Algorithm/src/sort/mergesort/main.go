package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generaterSlice(10)

	fmt.Println(slice)

	res := mergeSort(slice)

	fmt.Println(res)

}

func generaterSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func mergeSort(nums []int) []int {
	var length = len(nums)
	if length == 1 {
		return nums
	}

	var mid = int(length / 2)

	left := make([]int, mid)
	right := make([]int, length-mid)

	for i := 0; i < length; i++ {
		if i < mid {
			left[i] = nums[i]
		} else {
			right[i-mid] = nums[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	// leave left nums
	if i == len(left) {
		for j < len(right) {
			res[k] = right[j]
			k++
			j++
		}
	}
	// leave right nums
	if j == len(right) {
		for i < len(left) {
			res[k] = left[i]
			k++
			i++
		}
	}

	return res
}
