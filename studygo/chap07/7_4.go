package main

import (
	"time"
	"flag"
	"fmt"
)

var period = flag.Duration("period", 1 * time.Second, "sleep period")

func main()  {
	flag.Parse()
	fmt.Printf("sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}