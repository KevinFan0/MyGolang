package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	sendValue(t, ch)
	close(ch)	
}

// 递归向channel里传值
func sendValue(t *tree.Tree, ch chan int)  {
	if t != nil {
		sendValue(t.Left, ch)
		ch <- t.Value
		sendValue(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
// 使用写好的Walk函数来确定两个tree对象  是否一样 原理还是判断value值
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <- ch2 {		// ch1 关闭后   for循环自动跳出
			return false
		}
	}
	return true
}

func main() {
	// 打印 tree.New(1)的值
	var ch = make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
