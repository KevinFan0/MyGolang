package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	"os"
	"flag"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)


func main2() {
	symbol := [...]string{USD: "$", EUR:"€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
	s1 := os.Args[1]
	s2 := os.Args[2]
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	fmt.Println(Sha256Compare(&c1, &c2))
}

// 将[32]byte类型的数组清零
func zero(ptr *[32]byte)  {
	for i := range ptr{
		ptr[i] = 0
	}
}

// 更简洁
func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}

// 练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
func Sha256Compare(a, b *[32]byte) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count
}

// 练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。

func main()  {
	method := flag.String("method", "sha256", "select hash method(sha256,sha384,sha512)")
	text := flag.String("text", "", "input the string you want to hash")
	flag.Parse()
	switch *method {
	case "sha256":
		fmt.Println("%x\n", sha256.Sum256([]byte(*text)))
	case "sha384":
		fmt.Println("%x\n", sha512.Sum384([]byte(*text)))
	case "sha512":
		fmt.Println("%x\n", sha512.Sum512([]byte(*text)))
	default:
		panic("not support hash method")
	}
}