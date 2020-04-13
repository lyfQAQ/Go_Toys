package interface_test

import (
	"fmt"
	"testing"
)

// 接口 描述事物的外部行为 而非内部结构
type Shaper interface {
	Area() float32
}

// 该类型实现了 接口的方法
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func TestShaperInterface(t *testing.T) {
	sq1 := new(Square)
	sq1.side = 5

	var areaInterface Shaper
	// Square 实现了Shaper接口
	// 把 Squaer类型的变量赋值给Shaper类型的变量
	areaInterface = sq1

	t.Log("The square has area: ", areaInterface.Area())
}

type StockPositon struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s StockPositon) GetValue() float32 {
	return s.sharePrice * s.count
}

type Car struct {
	make  string
	model string
	price float32
}

func (c Car) GetValue() float32 {
	return c.price
}

// interface
type Valuable interface {
	GetValue() float32
}

func ShowValue(asset Valuable) {
	fmt.Print("Value of the asset is ", asset.GetValue(), "\n")
}

func TestValuableInterface(t *testing.T) {
	var v Valuable = StockPositon{"GOOG", 577.20, 4}
	ShowValue(v)

	v = Car{"BMW", "M3", 66500}
	ShowValue(v)
}
