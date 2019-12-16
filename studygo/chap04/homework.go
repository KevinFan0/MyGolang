package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	seen := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		// fmt.Println(input.Text())
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