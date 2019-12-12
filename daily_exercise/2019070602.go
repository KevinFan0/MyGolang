package main

import "fmt"

type X struct {
	a int
}

type Y struct {
	X
	b int
}

type Z struct {
	Y
	c int
}

func (x X) Print() {
	fmt.Printf("In X, a=%d\n", x.a)
}

func (x X) XPrint() {
	fmt.Printf("In X, a=%d\n", x.a)
}

func (y Y) Print() {
	fmt.Printf("In Y, b=%d\n", y.b)
}

func (z Z) Print() {
	fmt.Printf("In Z, c=%d\n", z.c)

	// 显示的完全路径调用内嵌字段的方法
	z.Y.Print()
	z.Y.X.Print()

}

func main2()  {
	x := X{a: 1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{
		Y: y,
		c: 3,
	}
	// z.a, z.Y.a, z.Y.X.a 三者是等价的， z.a, z.Y.a是z.Y.X.a的 简写
	println(z.a, z.Y.a, z.Y.X.a)

	//z = Z{}
	//z.a = 2
	//println(z.a, z.Y.a, z.Y.X.a)

	// 从外向内查找， 首先找到的是z的print()
	z.Print()
	// 从外向内查找， 最后找到的是x的xprint()
	z.XPrint()
	z.Y.XPrint()
}