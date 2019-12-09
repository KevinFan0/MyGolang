package main

import "fmt"

const boilingF = 212.0

func main_boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)

}

func incr(p *int) int {
	*p++	// 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	fmt.Println(&p)
	return *p
}

func main()  {
	v := 1
	incr(&v)	// side effect: v is now 2
	fmt.Println(incr(&v))	// "3" (and v is 3)
}
