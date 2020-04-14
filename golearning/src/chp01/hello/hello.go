package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello word")
	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("Hello World ", os.Args[1]) // 命令行参数
	}

	os.Exit(0) // 返回值
}
