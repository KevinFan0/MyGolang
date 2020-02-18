package main

import "fmt"

func main() {
	
	slice := []int{1,2,3,4,5}
	fmt.Printf("%p\n", &slice)
	modify(slice)
	fmt.Println(slice)
}


func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice1 := slice[:]
	// fmt.Printf("%p\n", &slice)
	// slice[1]=10
	slice1[2] = 20
}