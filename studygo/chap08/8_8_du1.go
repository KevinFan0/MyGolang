//创建一个程序来生成指定目录的硬盘使用情况报告，这个程序和Unix里的du工具比较相似

package main

import (
	"fmt"
	"path/filepath"
	"os"
	"io/ioutil"
	"flag"
)


func walkDir(dir string, fileSizes chan<- int64)  {
	for _, entry := range dirents(dir){
		// fmt.Println(entry.IsDir())
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}


// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}


func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	// Print the results.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDistUsage(nfiles, nbytes)
}

func printDistUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}