//创建一个程序来生成指定目录的硬盘使用情况报告，这个程序和Unix里的du工具比较相似

package main

import (
	"fmt"
	"sync"
	"time"
	"path/filepath"
	"os"
	"io/ioutil"
	"flag"
)


func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64)  {
	defer n.Done()
	for _, entry := range dirents(dir){
		// fmt.Println(entry.IsDir())
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}


// dirents returns the entries of directory dir.

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}							 // acquire token
	defer func() { <-sema }()					// release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// Print the results.
	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	loop:
		for {
			select {
			case size, ok := <-fileSizes:
				if !ok {
					break loop						// fileSizes was closed
				}
				nfiles++
				nbytes += size
			case <-tick:
				printDistUsage(nfiles, nbytes)
			}
		}
	printDistUsage(nfiles, nbytes)
}

func printDistUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}