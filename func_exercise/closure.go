// 闭包
// 闭包= 函数 + 引用环境


package main 

func fa(a int) func (i int) int {
	return func (i int) int {
		println(&a, a)
		a = a + 1
		return a
	}
} 

func main()  {
	f := fa(1)		//f引用的外部的闭包环境包括本次函数调用的形参a的值1
	g := fa(1)		//g引用的外部的闭包环境包括本次函数调用的形参a的值1
	
	// 此时f,g 引用的闭包环境中的a的值并不是同一个，而是两次函数调用产生的副本

	println(f(1))
	//多次调用f引用的是同一个副本a
	f(1)
	println(f(1))

	// g中a的值仍然是1

	println(g(1))
	println(g(1))
}

