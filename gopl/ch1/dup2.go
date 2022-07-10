package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	getFileName := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, getFileName, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // 打开文件
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, getFileName, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s\t%s\n", n, getFileName[line], line)
		}
	}

}

// map 作为函数参数时 传递的是引用
func countLines(f *os.File, counts map[string]int, getFileName map[string]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() { // 每次读一行 返回true 知道返回false
		counts[input.Text()]++
		getFileName[input.Text()] = fileName
	}
}
