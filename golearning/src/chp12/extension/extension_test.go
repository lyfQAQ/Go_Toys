package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet // 匿名成员
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Fan")
}
