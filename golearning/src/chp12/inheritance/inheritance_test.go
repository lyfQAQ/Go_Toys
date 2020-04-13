package inheritance

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	writeFirstProgram() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) writeFirstProgram() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) writeFirstProgram() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.writeFirstProgram())
	fmt.Printf("Type %T\n", p)
}

func TestPolymorphism(t *testing.T) {

	// 指针接收者实现 只能以指针方式使用，值接收者均可

	goProg := new(GoProgrammer)
	javaPro := new(JavaProgrammer)
	fmt.Printf("%T\n", goProg)
	writeFirstProgram(goProg)
	writeFirstProgram(javaPro)
}
