package main

import (
	"fmt"
	"os"
	"net/http"
	"golang.org/x/net/html"
)

func main()  {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

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

// 一个函数内部可以将另一个有多返回值的函数作为返回值，下面的例子展示了与findLinks有相同功能的函数
func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return fincLinks(url)
}

// 如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为bare return
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int){
	
}


func wordfreq(seen map[string]int) map[string]int {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		_, ok := seen[word]
		if ok {
			seen[word]++
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input:", err)
	}
	for k, v := range seen{
		fmt.Printf("%s\t%d\n", k, v)
	}
}