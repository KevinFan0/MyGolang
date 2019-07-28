package main

type INT int	// INT是一个使用预声明类型声明的自定义类型
type Map map[string]string		// Map是一个使用类型字面量声明的自定义类型
type myMap Map					// myMap是一个自定义类型Map声明的自定义类型
// INT，Map，myMap都是命名类型

// 自定义struct类型
// struct 初始化
type Person struct {
	name string
	age int
}

// （1）按照字段顺序进行初始化
a := Person{"andes", 18}
// （2）指定字段名进行初始化
a := Person{name: "andes", age: 18}
b := Person{
	name: "andes",
	age: 18,
}
c := Person{
	name: "andes",
	age: 18}
// （3）使用new创建内置函数，字段默认初始化为其类型的零值，返回值是指向结构的指针
p := new(Person)
//此时name为"", age是0
// （4）一次初始化一个字段 （不推荐）
p := Person{}
p.name = "andes"
p.age = 18
// 使用构造函数进行初始化

// 结构字段的特点

//标准库container/list

type Element struct {
	// 指向自身类型的指针
	next, prev *Element
	list *List
	Value interface()
}

// 自定义接口类型
// interface{}是接口字面量类型标识符， 所以i是非命名类型变量
// Reader 是自定义接口类型，属于命名类型
type Reader interface{
	Read(p []byte) (n int, err error)
}

// 类型方法
// 类型方法接收者是值类型
func (t TypeName)MethodName (ParaList) (Returnlist){
	// method body
}

// 类型方法接收者是指针
func (t *TypeName)MethodName(ParaList) (Returnlist)  {
	// method body
}

// t是接收者， 可以自由指定名称
// TypeName为命名类型的类型名
// MethodName为方法名，是一个自定义标识符
// ParamList是形参列表
// ReturnList是返回值列表

// 将类型的方法改写为常规函数
// 类型方法接收者是值类型
func TypeName_MethodName(t TypeName, otherParamList) (Returnlist) {
	//method body
}
// 类型方法接收者是指针
func TypeName_MethodName(t *TypeName, otherParamList) (Returnlist) {
	//method body
}

// 示例
type SliceInt []int
func (s SliceInt) Sum() int {
	sum := 0
	for _, i := range s{
		sum += i
	}
	return sum
}

//和上面方法等价
func SliceInt_Sum(s SliceInt) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

var s SliceInt = [] int{1,2,3,4}
s.Sum()
SliceInt_Sum(s)

// 方法调用
// 一般调用
type T struct{
	a int
}
func (t T) Get() int {
	return t.a
}

func (t *T) Set(i int)  {
	t.a = i
}

var t = &T{}
// 普通方法调用
t.Set(2)
// 普通方法调用
t.Get()

// 方法值ß

func (t *T) Print() {
	fmt.Printf("%p, %v, %d\n", t, t, t.a)
}

// method value
f := t.Set

//方法值调用
f(2)
t.Print()		// 2

//方法值调用
f(3)
t.Print()		// 3


