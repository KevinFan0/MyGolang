package main

import "fmt"


func main()  {
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
	fmt.Println(&p)
}


func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}