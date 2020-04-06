package main

import (
	"fmt"
	"math/rand"
	"math/cmplx"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func minus(x, y int) int{
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 /9
	y = sum -x
	return		// 无参return 返回已命名变量
}

// var可以在包内 或 函数内定义
var c, python, java bool

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("Hello,world")
	fmt.Println(add(rand.Intn(1), rand.Intn(10)))
	fmt.Println(math.Pi)
	fmt.Println(minus(2, 5))

	x, y := swap("hello", "world")
	fmt.Println(x, y)

	fmt.Println(split(17))

	// var函数内定义
	var i int
	fmt.Println(i, c, python, java)


	fmt.Printf("Type:%T  Value:%v ", ToBe, ToBe)
	fmt.Printf("Type:%T  Value:%v ", MaxInt, MaxInt)
	fmt.Printf("Type:%T  Value:%v ", x, x)
}
