package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a Integer) Add(b Integer) {
	a += b
}

func (a *Integer) Add2(b Integer) {
	*a += b
}

//定义结构体
type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

//结构体的初始化
rect1 := new(Rect)
rect2 := &Rect{}
rect3 := &Rect{0, 0, 100, 200}
rect4 := &Rect{width: 100, height: 200}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
	a.Add(2)
	fmt.Println("a =", a)
	a.Add2(2)
	fmt.Println("a =", a)
}
