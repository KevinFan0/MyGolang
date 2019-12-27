package main

import (
    "fmt"
    "os"
)

func squares() func()	int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main()  {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// 要求首先创建一些目录，再将目录删除。在下面的例子中我们用函数值来完成删除操作
func catchdiedai() {
    var rmdis [] func()
    for _, d := range tempDirs() {
        dir := d
        os.MkdirAll(dir, 0755)
        rmdirs = append(rmdirs, func() {
            os.RemoveAll(dir)
        })
    }
    for _, r := range rmdirs {
        rmdir()
    }
}