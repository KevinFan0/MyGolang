package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	
)

type WordsCounter int

func (c *WordsCounter) Write(w []string) (int, error) {
	*c += WordsCounter(len(w))
	return len(w), nil
}



func main()  {
	var c WordsCounter
	// c.Write([]string{"hello", "world"})
	// fmt.Println(c)
	wordlist := []string{}
	input := []string{"hello", "world", "\n"} 
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner, err = scanner.ReadString('\n')
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	fmt.Println(c.Write(wordlist))
}