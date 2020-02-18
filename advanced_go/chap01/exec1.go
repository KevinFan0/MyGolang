package main

import "fmt"

var a1 [3]int                    // 定义长度为3的int型数组, 元素全部为0
var b1 = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6



var a = [...]int{1, 2, 3}
var b = &a							//b是指向数组的指针

func main()  {
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	for i, v := range b {
		fmt.Println(i, v)
	}
}

// 其他数组类型
//字符串数组
var s1 = [2]string{"hello", "world"}
var s2 = [...]string{"你好", "世界"}
var s3 = [...]string{1: "世界", 2: "你好", }


//结构体数组
var line1 [2]image.Point
var line2 [...]imate.Point({imate.Point{X: 0, Y: 0}, image.Point{X:1, Y:1}})
var line3 [...]image.Point{{0, 0}, {1, 1}}


//接口数组
var unknown1 [2]interface{}
var unknown2 [...]interface{(123, "你好")}

//管道数组
var chanList = [2]chan int{}

//空数组 定义一个长度为0的数组
var d [0]int
var e = [0]int{}
var f = {...}int{}


c1 := make(chan [0]int)
go func() {
	fmt.Println("c1")
	c1 <- [0]int{}
}()
<- c1

c2 := make(chan struct{})
go func(){
	fmt.Println("c2")
	c2 <- struct{}{}	//struct{}部分是类型， {}表示对应的结构体值
}()
<- c2