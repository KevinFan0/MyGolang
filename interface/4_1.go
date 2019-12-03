package _interface

// 接口声明
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Writer(p []byte) (n int, err error)
}

// 如下3种声明是等价，最终展开模式都是第三种格式

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriter interface {
	Reader
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 声明新接口类型的特点
// (1) 接口的命名一般以er 结尾
// (2) 接口定义的内部方法声明不需要func引导
// (3) 在接口定义中， 只有方法声明没有方法实现