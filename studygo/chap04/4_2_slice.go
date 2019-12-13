package main
 import (
	 "fmt"
	 "unicode"
 )

 // reverse函数在原内存空间将[]int类型的slice反转，而且它可以用于任意长度的slice
 func reverse(s []int)  {
	 for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		 s[i], s[j] = s[j], s[i]
	 }
 }

 //比较两个slice类型是否全部相等
 func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// 专门用于处理[]int类型的slice
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x){
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// 给定一个字符串列表，下面的nonempty函数将在原有slice内存空间之上返回不包含空字符串的列表
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 要删除slice中间的某个元素并保存原有的元素顺序，可以通过内置的copy函数将后面的子slice向前依次移动一位完成：
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

//重写reverse函数，使用数组指针代替slice
func reversepoint(p *[]int)  {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

// 编写一个rotate函数，通过一次循环完成旋转。
func rotate(slice []int, n int) []int {
	if n <= len(slice) {
		n = n 
	}else {
		n = n % len(slice)
	}
	slice = append(slice, slice[:n]...) // push v
	slice = slice[n:]
	return slice

}

//写一个函数在原地完成消除[]string中相邻重复的字符串的操作
func RemoveRepeatChar(s []string) []string {
	if len(s) < 1{
		return s
	}
	current = s[0]
	index := 0
	for i := 1; i < len(s); i++ {
		if s[i] != current{
			index += 1
			s[index] = s[i]
			current = s[index]
		}
	}
	return s[:index]
}

//编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func RemoveByteSpace(r []rune)  {
	index := 0
	flag :=  unicode.IsSpace(r[0])
	for i := 1; i < len(r); i++ {
		if unicode.IsSpace(r[i]){	
			if flag {
				flag = false
				index += 1
				r[index] = r[i]
			}
		}else {
			index += 1
			r[index] = r[i]
		}
	}

}


 func main()  {
	 a := [...]int{0, 1, 2, 3, 4, 5}
	//  reversepoint(&a)
	//  fmt.Println(a)
	 /* 一种将slice元素循环向左旋转n个元素的方法是三次调用reverse反转函数，第一次是反转开头的n个元素，然后是反转剩下的元素，最后是反转整个slice的元素。（如果是向右循环旋转，则将第三个函数调用移到第一个调用位置就可以了。）*/
	 // 循环向左旋转2个元素
	//  reverse(a[:2])
	//  reverse(a[2:])
	//  reverse(a[:])
	//  fmt.Println(a) // "[2 3 4 5 0 1]"
	 fmt.Println(rotate(a[:], 6))
 }

