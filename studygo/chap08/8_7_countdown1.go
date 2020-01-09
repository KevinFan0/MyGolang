package main

import (
	"fmt"
	"time"
	"os"
	
)

func main()  {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte,1))			// read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort")
	tick := time.Tick(1 * time.Second)
	for countdown :=10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <- tick:
			// Do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	/*
	countdown2:
		select {
		case <- time.After(10 * time.Second):
			// do nothing.
			{}
		case <- abort:
			fmt.Println("Launch aborted!")
			return
		}
	*/
	/*
	countdown1:
		tick := time.Tick(1 * time.Second)
		for countdown := 10; countdown > 0; countdown-- {
			fmt.Println(countdown)
			<- tick
		}	
	*/
	launch()
}


ch := make(chan int, 1)
for i:=0; i < 10; i++ {
	select {
	case x := <-ch:
		fmt.Println(x)			// "0" "2" "4" "6" "8"
	case ch <- i:
	}
}
