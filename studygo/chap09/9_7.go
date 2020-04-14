package main

import (
	"net/http"
	"fmt"
	"studygo/chap09/memo"
	"time"
	"io/ioutil"
	"log"
	"sync"
)

// 这个函数会去进行HTTP GET请求并且获取http响应body。对这个函数的调用本身开销是比较大的，所以我们尽量避免在不必要的时候反复调用
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)				// ReadAll会返回两个结果，一个[]byte数组和一个错误，不过这两个对象可以被赋值给httpGetBody的返回声明里的interface{}和error类型
}

func incomingURLs() []string {
    return []string{"https://www.segmentfault.com", "https://www.baidu.com", "http://www.sogou.com", "http://che.sogou.com", "http://m.che.sogou.com", "https://www.segmentfault.com", "https://www.baidu.com", "https://www.sogou.com", "http://che.sogou.com", "http://m.che.sogou.com", "https://www.segmentfault.com", "https://www.baidu.com", "https://www.sogou.com", "http://che.sogou.com", "http://m.che.sogou.com"}
}

func main()  {
	m := memo.New(httpGetBody)
	// 使用sync.WaitGroup来等待所有的请求都完成之后再返回
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}