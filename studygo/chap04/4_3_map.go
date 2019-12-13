package main

import (
	"fmt"
	"sort"
)


func MapSortKey() {
	var names []string
	var ages []int
	names := make([]string, 0, len(ags))
	for name := range ages {
		names = append(names, name)
	}
	
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y){
		return false
	}
	for k, xv:= range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}


func main()  {
	
}