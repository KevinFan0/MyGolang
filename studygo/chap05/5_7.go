package main

import (
	"fmt"
	"os"
)

//errorf函数构造了一个以行号开头的，经过格式化的错误信息, interfac{}表示函数的最后一个参数可以接收任意类型
func errorf(linenum int, format string, args ...interface{})  {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr,)
}

func main()  {
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}