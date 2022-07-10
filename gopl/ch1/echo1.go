// print cmd
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	for idx, arg := range os.Args {
		fmt.Printf("%d %s\n", idx, arg)
	}

}
