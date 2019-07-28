// 4.2 接口运算
// 类型断言

package main
import "fmt"


type Inter interface{
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	// Name string
}

func (St) Ping()  {
	println("ping")
}

func (*St) Pang()  {
	println("pang")
}

// func (St) String()  {
// 	println("string")
// }


func main()  {
	// st := &St{"andes"}
	// var i interface{} = st
	// 判断i绑定的实例是否实现了接口类型Inter
	// o := i.(Inter)
	// o.Ping()
	// o.Pang()
	// 如下语句会引发panic，因为i没有实现接口Anter(没有实现String方法, 只要定义一个string方法，就不会报错)
	// p := i.(Anter)
	// p.String()

	// 判断i绑定的实例是否就是具体类型St
	// s := i.(*St)
	// fmt.Printf("%s\n", s.Name)

	// 利用 comma，ok表达式
	// 判断i绑定的实例是否实现了接口类型Inter
	// if o, ok := i.(Inter); ok{
	// 	o.Ping()
	// 	o.Pang()
	// }

	// if p, ok := i.(Anter); ok{
	// 	// i 没有实现接口Anter，所以程序不会执行到这里
	// 	p.String()
	// }

	// // 判断i绑定的实例是否就是具体类型st
	// if s, ok := i.(*St); ok {
	// 	// fmt.Printf("%s\n", s.Name)
	// }


	// // 类型查询
	// var i io.Reader
	// switch v := i.(type) {		// 此处i 是未初始化的接口变量，所以v为nil
	// case nil:
	// 	fmt.Printf("%T\n", v)	//<nil>
	// default:
	// 	fmt.Printf("default")
	// }

	// // 如果case后面是一个接口类型名，且接口变量i绑定的实例类型实现了该接口类型的方法，则匹配成功，v的类型是接口类型，v底层绑定的实例是i绑定具体类型实例的副本
	// f, err := os.OpenFile("xxxx.txt", os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// var i io.Reader = f
	// switch v := i.(type) {
	// // i的绑定实例是*osFile类型，实现了io.ReadWrtier接口，所以case匹配成功
	// case io.ReadWrtier:
	// 	// v是io.ReadWrtier接口类型，所以可以调用Write方法
	// 	v.Write([]byte("io.ReadWrtier\n"))
	// // 由于上一个case已经匹配，就算这个case也匹配，也不会走到这里
	// case *os.File:
	// 	v.Write([]byte("*os.File\n"))
	// 	v.Sync()
	// default:
	// 	return
	// }
	// // 如果case后面就算一个具体类型名，且接口变量i绑定的实例类型和该具体类型相同，则匹配成功，此时v就是该具体类型变量，v的值是i绑定的实例值的副本
	// var i io.Reader = f
	// switch v := i.(type) {
	// // 匹配成功，v的类型就是具体类型*os.File
	// case *os.File:
	// 	v.Write([]byte("*os.File\n"))
	// 	v.Sync()
	// // 由于上一个case已经匹配，就算这个case也匹配，也不会走到这里
	// case io.ReadWrtier:
	// 	v.Write([]byte("io.ReadWrtier\n"))
	
	
	// default:
	// 	return
	// }


	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)

	if it != nil {
		it.Pang()
		// 下面的语句会导致panic
		// 方法转换为函数调用，第一个参数是st类型，由于*St是nil，无法获取指针所指的对像值，所以导致panic
		it.Ping()
	}
}

// 4.3 空接口
// 空接口和泛型
// 如果一个函数需要接受任意类型的参数，则参数类型可以使用空接口类型，这是弥补没有泛型的一种手段
// 典型的就是fmt标准包里的print函数
// func Fprint(w io.Writer, a ...interface{}) (n int, err error)





