package main

import (
	"log"
	"fmt"
	"os"
	"studygo/chap08/links"
)

var tokens = make(chan struct{}, 20)
// 将对links.Extract的调用操作用获取、释放token的操作包裹起来，来确保同一时间对其只有20个调用。信号量数量和其能操作的IO资源数量应保持接近。
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}			// acquire a token
	list, err := links.Extract(url)		
	<- tokens						// release the token
	if err != nil {
		log.Print(err)
	}
	return list
}


func main()  {
	worklist := make(chan []string)
	var n int									// number of pending sends to worklist
	 // Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()
	 // Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
	
}