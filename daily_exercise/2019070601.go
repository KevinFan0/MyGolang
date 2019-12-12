// 方法集
package main

import "fmt"

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


// 值调用和表达式调用的方法集
//(1) 通过类型字面量显示地进行值调用和表达式调用，可以看到这种情况下编译器不会做自动转换
type Data struct {}

func (Data) TestValue() {}

func (*Data) TestPointer()  {}

// 这种字面量显示调用，无论值调用还是表达式调用，编译器都不会进行方法集的自动转换
// *Data 方法集是TestPointer和TestValue；Data 方法集只有TestValue
(*Data) (&struct{}{}).TestPointer()		// 显示的调用
(*Data)(&struct{}{}).TestValue()		// 显示的调用

(Data) (struct{}{}).TestValue()			// method value
Data.TestValue(struct{}{})				// method expression

// 如下调用因为方法集和不匹配而失败
// Data.TestPointer(struct{}{})			// type Data has no method TestPointer
// (Data)(struct{}{}).TestPointer()		// cannot call pointer method on Data(struct{} literal)


//(2) 通过类型变量进行值调用和表达式调用，使用值调用方式调用时编译器会进行自动转换，使用表达式调用方式调用时编译器不会进行自动
// 转换，会进行严格的方法集检查
// 声明一个类型变量a
var a Data = struct{}{}

//表达式调用编译器不会进行自动转换
Data.TestValue(a)
//Data.TestValue(&a)
(*Data).TestPointer(&a)
// Data.TestPointer(&a)		// type Data has no method TestPointer

// 值调用编译器会进行自动转换

f := a.TestValue
f()

y := (&a).TestValue		//编译器帮助转换a.TestValue
y()

g := a.TestPointer		// 会转换为(&a).TestPointer
g()

x := (&a).TestPointer
x()

func main3()  {
	var a Int = 10
	var  b Int = 20
	c := a.Max(b)
	c.Print()		//value=20
	(&c).Print()	//value=20 内部被编译器转换为c.Print()
	a.Set(20)		//内部被编译器转化为（&a).Set(20)
	a.Print()		//value=20
	(&a).Set(30)
	a.Print()		//value=30
}