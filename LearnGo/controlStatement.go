package main

import (
	"fmt"
	"math"
	"runtime"
)

// if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if条件前 可以加一个语句
// 变量v 在if else 块 中有效
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	sum := 0

	// for
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	n := 1
	for n < 1000 {
		n += n
	}
	fmt.Println(n)

	// while
	m := 1
	for m < 1000 {
		m += m
	}
	fmt.Println(m)

	// if
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	
	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	defer fmt.Println("world")

	fmt.Println("hello")
}
