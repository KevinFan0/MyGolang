// recover 
// recover只有在defer后面的函数体内被直接调用才能捕获panic终止异常，否则返回nil，异常继续向上传递

// 这个捕获失败
/*
defer recover()

// 这个捕获失败
defer fmt.Println(recover())

// 这个嵌套两层也会捕获失败
defer func ()  {
	func ()  {
		Println("defer inner")
		recover()		//无效
	}()
}()
*/

package main
import (
	"fmt"
	"time"
)
// 如下场景会捕获成功
/*
defer func ()  {
	Println("defer inner")
	recover()
}()
*/
func except()  {
	recover()
}

func test()  {
	defer except()
	panic("test panic")
}

// 函数不能捕获内部新启动的goroutine所抛出的panic
func do()  {
	// 这里并不能捕获da函数中的panic
	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	go da()
	go db()
	time.Sleep(3 * time.Second)
}

func da(){
	panic("panic da")
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}

func db(){
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}
func main3()  {
	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	// 只有最后一次panic调用能够捕获
	defer func ()  {
		panic("first defer panic")
	}()
	defer func ()  {
		panic("second defer panic")
	}()
	panic("main body panic")
}
