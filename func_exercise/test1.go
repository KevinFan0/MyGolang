package main

import (
	"fmt"
)
func chvalue(a int) int {
	a = a + 1
	return a
}

func chpointer (a *int) {
	*a = *a + 1
	return
}

func main1()  {
	a := 10
	chvalue(a)				//实参传递给形参是值拷贝
	fmt.Println(a)
	chpointer(&a)			//实参传递给形参仍然是值拷贝，只不过复制的是a的地址值
	fmt.Println(a)
}


func sum(a, b  int) int {
	return a + b
}

func main2()  {
	sum(3, 4)		//直接调用
	f := sum		// 有名函数可以直接赋值给变量
	f(1, 2)
}


// 匿名函数
//匿名函数被直接赋值函数变量
var sum1 = func (a, b int) int {
	return a + b
}

func doinput(f func(int, int) int, a, b int) int {
	return f(a, b)
}

// 匿名函数作为返回值
func wrap(op string) func (int, int) int {
	switch op {
	case "add":
		return func (a, b int) int {
			return a + b
		}
	case "sub":
		return func (a, b int) int {
			return a - b
		}
	default:
		return nil
	}
} 

func main3()  {
	// 匿名函数直接被调用
	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	sum(1, 2)
	// 匿名函数作为实参
	doinput(func (x, y int) int {
		return x + y
	}, 1, 2)
	
	opFunc := wrap("add")
	re := opFunc(2, 3)
	fmt.Printf("%d\n", re)
}


// defer 用法
func main()  {
	//先进后出
	defer func ()  {
		println("first")
	}()
	defer func ()  {
		println("second")
	}()
	println("function body")
}
// 结果是先注册后执行


//如果defer 语句位于return语句之后，则defer没有注册，不会执行

//主动调用os.Exits(int)退出进程时，defer将不再被执行

// defer 语句位置不当，有可能导致panic， 一般defer语句放在错误检查语句之后

// defer中最好不要对有名返回值参数进行操作，否则会引发匪夷所思对结果