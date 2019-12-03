// 组合的方法集
// 遵守的规则
// (1) 若类型S包含匿名字段T， 则S的方法集包含T的方法集
// (2) 若类型S包含匿名字段*T， 则S的方法集包含T和*T的方法集
// (3) 不管类型S中嵌入的匿名字段是T还是*T，*s方法集总是包含T和*T方法集


package main

type X struct {
	a int
}

type Y struct {
	X
}

type Z struct {
	*X
}

func (x X) Get() int {
	return x.a
}

func (x *X) Set(i int) {
	x.a = i
}

func main1()  {
	x := X{a: 1}
	y := Y{
		X: x,
	}
	println(y.Get())		//1
	// 此处编译器做了自动转换
	y.Set(2)
	println(y.Get())		//2
	// 为了不让编译器做自动转换，使用方法表达式调用方式
	// Y内嵌字段x， 所以type Y的方法集是Get，type *Y的方法集是Set Get
	(*Y).Set(&y, 3)
	// type Y的方法集合并没有Set方法，所以下一句编译不能通过
	// Y.Set(y,3)
	println(y.Get())		//3
	z := Z{
		X: &x,
	}
	// 按照嵌套字段的方法集的规则
	// Z 内嵌字段*X，所以type z 和type *Z 方法集都包含类型X定义的方法Get和Set

	// 为了不让编译器做自动转换，仍然使用方法表达式调用方式
	Z.Set(z, 4)
	println(z.Get())		//4
	(*Z).Set(&z, 5)
	println(z.Get())		//5
}