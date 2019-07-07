package main

type INT int	// INT是一个使用预声明类型声明的自定义类型
type Map map[string]string		// Map是一个使用类型字面量声明的自定义类型
type myMap Map					// myMap是一个自定义类型Map声明的自定义类型
// INT，Map，myMap都是命名类型

// 自定义struct类型
// struct 初始化
type Person struct {
	name string
	age int
}

// （1）按照字段顺序进行初始化
a := Person{"andes", 18}
// （2）指定字段名进行初始化
a := Person{name: "andes", age: 18}
b := Person{
	name: "andes",
	age: 18,
}
c := Person{
	name: "andes",
	age: 18}
// （3）使用new创建内置函数，字段默认初始化为其类型的零值，返回值是指向结构的指针
p := new(Person)
//此时name为"", age是0
// （4）一次初始化一个字段 （不推荐）
p := Person{}
p.name = "andes"
p.age = 18
// 使用构造函数进行初始化

