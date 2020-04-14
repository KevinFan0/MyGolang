package main

import (
	"fmt"
	"studygo/chap09/bank"
)


func main()  {
	// Alice
	go func ()  {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()
	// Bob
	go bank.Deposit(100)
	
}


var icons = make(map[string]image.Image)
func loadIcon(name string) image.Image 

// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}


type Cake struct { state string }

func baker(cooked chan<- *Cake)  {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake						// baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake						// icer never touches this cake again
	}
}