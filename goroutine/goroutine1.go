package main

import (
	"runtime"
	"time"
)

// func main() {
// 	go func() {
// 		sum := 0
// 		for i := 0; i < 10000; i++ {
// 			sum += i
// 		}
// 		println(sum)
// 		time.Sleep(1 * time.Second)
// 	}()
// 	// NumGoroutin 可以返回当前程序的goroutine数目
// 	println("NumGoroutine=", runtime.NumGoroutine())

// 	// main goroutine故意sleep 5秒，防止其提前退出
// 	time.Sleep(5 * time.Second)
// }

func sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	time.Sleep(1 * time.Second)
	second := time.Second
	println(second)
}

func main() {
	go sum()
	// NumGoroutin 可以返回当前程序的goroutine数目
	println("NumGoroutine=", runtime.NumGoroutine())
	// main goroutine故意sleep 5秒，防止其提前退出
	time.Sleep(5 * time.Second)
}
