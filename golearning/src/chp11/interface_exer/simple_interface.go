package main

import "fmt"

type simpler interface {
	Get() int
	Set(v int)
}

type simple struct {
	value int
}

func (s *simple) Get() int {
	return s.value
}

func (s *simple) Set(v int) {
	s.value = v
}

func testInterface(sr simpler) {
	fmt.Printf("value is %d\n", sr.Get())
	fmt.Printf("Set value 5\n")
	sr.Set(5)
	fmt.Printf("value is %d\n", sr.Get())
}

func main() {
	var sr simpler

	sr = new(simple)
	testInterface(sr)
}
