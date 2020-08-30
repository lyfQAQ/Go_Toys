package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	for idx, s := range os.Args {
		fmt.Println(idx, s)
	}
}
