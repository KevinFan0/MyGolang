//代码清单2-3 map1.go
package main
import "fmt"

type PersonInfo struct{
	ID string
	Name string
	Address string
}

func main3(){
	var personDB map[string] PersonInfo		// 变量声明
	personDB = make(map[string] PersonInfo)		// 创建

	//往map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 103,..."}

	// 从这个map里查找键为"1234"的信息
	person, ok := personDB["12345"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 12345.")
	} else {
		fmt.Println("Did not find person with ID 12345.")
	}
}