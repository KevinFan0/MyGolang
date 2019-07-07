package main

import "fmt"

// 使用type声明的命名类型
type Person struct{
	name string
	age int
}

func main2()  {
	// 使用struct字面量声明的是未命名类型
	a := struct{
		name string
		age int
	}{"andes", 18}

	fmt.Printf("%T\n", a)	//struct {name string; age int}
	fmt.Printf("%v\n", a)	//{andes 18}

	b := Person{"tom", 21}
	fmt.Printf("%T\n", b)	//main.Person
	fmt.Printf("%v\n", b)	//{tom 21}
}


// 类型直接赋值
// a是类型为T1的变量，或者a本身就是一个字面常量或nil
// 如果如下语句可以执行，则称之为类型T1可以赋值给类型T2
// var b T2 = a

// 需满足条件：
// （1） T1和T2的类型相同
// （2） T1和T2具有相同的底层类型，并且T1和T2里面至少有一个是未命名类型
// （3） T2是接口类型，T1是具体类型， T1的方法集是T2方法集的超集
// （4） T1和T2是通道类型，它们拥有相同的元素类型，并且T1和T2中至少有一个是未命名类型
// （5） a是预声明标识符nil，T2是pointer, function, slice, map, channel, interface类型中的一个
// （6） a是一个字面常量值，可以用来表示类型T的值


type Map map[string]string

func (m Map) Print() {
	for _, key := range m{
		fmt.Println(key)
	}
}

type iMap Map

// 只要底层类型是slice, map等支持range的类型字面量，新类型仍然可以使用range迭代
func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main3() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"

	// mp与ma有相同的底层类型map[string]string, 并且mp是未命名类型
	// 所以mp可以直接赋值给ma
	var ma Map = mp 

	// im与ma虽然有相同的底层类型map[string][string], 但它们中没有一个是未命名类型
	// 不能赋值，如下语句不能通过编译
	// var im IMap = ma

	ma.Print()
	//im.Print()

	// Map实现了Print(), 所以其可以赋值给接口类型变量

	var i interface {
		Print()
	} = ma

	i.Print()
	s1 := [] int{1,2,3}
	var s2 slice
	s2 = s1
	s2.Print()

}


// 类型强制转换
// 非常量类型的变量x可以强制转化并传递给类型T，需要满足任何如下条件：
// （1）x可以直接赋值给T类型变量
// （2）x的类型和T具有相同的底层类型
// im与ma虽然有相同的底层类型map[string][string], 但它们中没有一个是字面量类型， 不能直接赋值，可以强制进行类型转化
//  var im iMap = ma 				(no)
// var im iMap = (iMap) (ma)		(yes)
// （3）x的类型和T都是未命名的指针类型，并且指针指向的类型具有相同的底层类型
// （4）x的类型和T都是整型，或者都是浮点型
// （5）x的类型和T都是复数类型
// （6）x是整数值或[]byte类型的值， T是string类型
// （7）x是一个字符串，T是[]byte或[] rune

// 字符串和字符切片之间的转换
// s := "hello, world"
// var s []byte
// a = []byte(s)
// var b = string
// b = string(a)

// 注意：
// （1）数值类型和string类型之间的相互转换可能造成值部分丢失。string和数字之间的转换可以使用标准库strconv
// （2）go语言没有语言机制支持指针和interger之间的直接转换，可以使用标准库中的unsafe包进行处理
