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

// 使用map来记录提交相同的字符串列表的次数。它使用了fmt.Sprintf函数将字符串列表转换为一个字符串以用于map的key，通过%q参数忠实地记录每个字符串元素的信息
var m = make(map[string]int)
func k(list []string) string { return fmt.Sprintf("%q", list) }
func Add(list []string)  { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main()  {
	
}