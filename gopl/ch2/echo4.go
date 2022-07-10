//-n=true -s=@ 12 34 56
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "new line")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	//12@34@56
	if !*n {
		fmt.Println()
	}
}
