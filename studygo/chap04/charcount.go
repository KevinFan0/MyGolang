package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)


// 下面的程序用于统计输入中每个Unicode码点出现的次数
func main()  {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid, letter := 0, 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		// fmt.Println(n)
		if err == io.EOF{
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r){
			letter++
		}
		if r == unicode.ReplacementChar && n == 1{
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	
	fmt.Printf("letter: %d\n", letter)
	fmt.Printf("rune\tcount\n")
	for c, n := range counts{
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0{
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}


// 图graph的key类型是一个字符串，value类型map[string]bool代表一个字符串集合。从概念上讲，graph将一个字符串类型的key映射到一组相关的字符串集合，它们指向新的graph的key
var graph = make(map[string]map[string]bool)

func addEdge(from, to string)  {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}