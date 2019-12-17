package main

import (
	"math"
	"fmt"
	"os"
	// "studygo/chap05/func5_1"
	"golang.org/x/net/html"
	// "reflect"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}


// visit函数遍历HTML的节点树，从每一个anchor元素的href属性获得link,将这些links存入字符串数组中，并返回这个字符串数组。
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a"{
		for _, a := range n.Attr{
			if a.Key == "href"{
				// fmt.Println(reflect.TypeOf(a))
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }
	return links
}


// 在函数outline中，我们通过递归的方式遍历整个HTML结点树，并输出树的结构。在outline内部，每遇到一个HTML元素标签，就将其入栈，并输出。
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}


func main()  {
	// fmt.Println(hypot(3, 4))
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	// for _, link := range visit(nil, doc) {
		// fmt.Println(link)
	// }
	// outline(nil, doc)
	// m := make(map[string]int)
	// m = visitCount(m, doc)
	// for item, count := range m {
	// 	fmt.Printf("%s\t%d\n", item, count)
	// }
	fmt.Println(visit2(nil, doc))
}

// visit函数遍历HTML的节点树，从每一个anchor元素的href属性获得link,将这些links存入字符串数组中，并返回这个字符串数组。
func visitCount(m map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
			m[n.Data]++
		}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitCount(m, c)
	}
	return m
	// return visit(visit(m, n.FirstChild), n.NextSibling)
}


func visit2(links []string, n *html.Node) []string {
	if n != nil && n.Type == html.TextNode {
		// fmt.Println(n.Attr)
		for _, a := range n.Attr{
			// fmt.Println(a)
			if a.Key != "script" && a.Key != "style"{
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.TextNode{
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit2(links, c)
	}
	return links
}