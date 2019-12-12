package main

// 方法集

import (
	"fmt"
	"time"
)

type Int int

func (a Int) Max(b Int) Int {
	if a >= b{
		return a
	} else {
		return b
	}
}

func (i * Int) Set(a Int) {
	*i = a
}

func (i Int) Print() {
	fmt.Printf("value=%d\n", i)
}

func main()  {
	var a Int = 10
	var  b Int = 20
	c := a.Max(b)
	c.Print()		//value=20
	(&c).Print()	//value=20 内部被编译器转换为c.Print()
	a.Set(20)		//内部被编译器转化为（&a).Set(20)
	a.Print()		//value=20
	(&a).Set(30)
	a.Print()		//value=30
	var d string = "abcdefg"
	//var e float32 = 0.0001
	fmt.Println(string(d)[:2])
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.String())
}