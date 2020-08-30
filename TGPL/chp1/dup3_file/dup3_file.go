package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var res []string
	for _, file := range os.Args[1:] {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadFile: %v\n", nil)
		}
		counts := make(map[string]int)
		for _, line := range strings.Split(string(content), "\n") {
			counts[line]++
		}
		for _, n := range counts {
			if n > 1 {
				res = append(res, file)
				break
			}
		}
	}
	fmt.Println(res)
}
