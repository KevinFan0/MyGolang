package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"net/http"
)

func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }


func main()  {
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)
	url := os.Args[1:]
	resp, err := http.Get(url[0])
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	forEachNode(doc, startElement, endElement)
}


// 5.2节的findLinks函数使用了辅助函数visit,遍历和操作了HTML页面的所有结点。使用函数值，我们可以将遍历结点的逻辑和操作结点的逻辑分离，使得我们可以复用遍历的逻辑，从而对结点进行不同的操作。
// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node))  {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// 现在我们有startElemen和endElement两个函数用于输出HTML元素的开始标签和结束标签<b>...</b>
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode{
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
	if n.Type == html.CommentNode{
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node)  {
	if n.Type == html.ElementNode{
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}