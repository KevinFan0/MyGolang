package main

import (
	"fmt"
	"bytes"
	"strings"
	"os"
	"sort"
	"reflect"
)

func main() {
	// fmt.Println(basename("a/b/c.go"))
	// fmt.Println(basename("c.d.go"))
	// fmt.Println(basename("abc"))
	// fmt.Println(basename2("a/b/c.go"))
	// fmt.Println(basename2("c.d.go"))
	// fmt.Println(basename2("abc"))
	// fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
	// fmt.Println(comma2("12345"))
	// s := "12345"
	// n := len(s) % 3
	// fmt.Println(n3
	// for i := 1; i < len(os.Args); i++ {
		// fmt.Println(comma3(os.Args[i]))
	// }
	s1 := os.Args[1]
	s2 := os.Args[2]
	// s2 := strings.Split("")
	// s2arr := []byte(s2)
	// sort.Sort(sort.StringSlice(s1arr))
	fmt.Println(isrepeat(s1, s2))
}

// basename()将看起来像是系统路径的前缀删除，同时将看似文件类型的后缀名部分删除
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "/" {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "." {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//函数的功能是将一个表示整值的字符串，每隔三个字符插入一个逗号分隔符，例如“12345”处理后成为"12,345"
func comma(s string) string {
	n := len(s)
	if n <= 3{
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])
}

func intsToString(values []int) string { 
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// "12345"
func comma2(s string) string {
	
	var buf bytes.Buffer
	n := len(s) % 3
	if n == 0{
		n = 3
	}
	buf.WriteString(s[:n])
	for n < len(s) {
		buf.WriteByte(',')
		buf.WriteString(s[n:n+3])
		n += 3
	}
	return buf.String()
}


func comma3(s string) string {
	
	var buf bytes.Buffer
	s2 := strings.Split(s, ".")
	// 先算整数部分
	n := len(s2[0]) % 3
	if n == 0{
		n = 3
	}
	buf.WriteString(s2[0][:n])
	for n < len(s2[0]) {
		buf.WriteByte(',')
		buf.WriteString(s2[0][n:n+3])
		n += 3
	}
	// 再算小数部分
	n2 := len(s2[1]) % 3
	if n2 == 0{
		n2 = 3
	}
	buf.WriteString(s2[1][:n])
	for n < len(s2[0]) {
		buf.WriteByte(',')
		buf.WriteString(s2[0][n:n+3])
		n += 3
	}
	
	return buf.String()
}

//编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func isrepeat(s1, s2 string) bool {
	// 先声明一个未指定大小的数组来定义切片
	var slice1 []string
	var slice2 []string
	for i := 0; i < len(s1); i++ {
		slice1 = append(slice1, string(s1[i]))
	}
	for i := 0; i < len(s2); i++ {
		slice2 = append(slice2, string(s2[i]))
	}
	sort.Strings(slice1)
	sort.Strings(slice2)
	s1 = strings.Join(slice1, "")
	s2 = strings.Join(slice2, "")
	return reflect.DeepEqual(s1, s2)
}