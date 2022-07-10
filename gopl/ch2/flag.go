//通过命令行参数给变量赋初值 -n=true -s="123"
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 返回*bool
	n := flag.Bool("n", false, "n bool")
	// 返回*string
	s := flag.String("s", " ", "seq")

	// 在flag参数定义之后访问之前 执行解析
	flag.Parse()

	fmt.Println("n value", *n, "s value", *s)

}
