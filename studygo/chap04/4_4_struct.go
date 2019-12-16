package main

import (
	"time"
	"fmt"
)
type Employee struct {
	ID			int
	Name		string
	Address		string
	DoB			time.Time
	Position	string
	Salary		int
	ManagerID	int
}

type Point struct { X, Y int }

var dilbert Employee

func EmployeeByID(id int) *Employee { return &dilbert}


// 利用二叉树来实现一个插入排序

type tree struct{
	value			int
	left, right		*tree
}

// Sort sorts values in place
func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	res := appendValues(values[:0], root)
	return res
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	}else {
		t.right = add(t.right, value)
	}
	return t
}


//结构体可以作为函数的参数和返回值
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// 如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent /100
}

// 如果要在函数内部修改结构体成员的话，用指针传入是必须的；
func AwardAnnualRaise(e *Employee)  {
	e.Salary = e.Salary * 105 / 100
}

// 比较两个结构体
func equalStruct()  {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)		// false
	fmt.Println(p == q)							// false
}

type address struct {
	hostname 	string
	port		int
}

func StructMapKey() {
	hits := make(map[address]int)
	hits[address{"golang.org", 433}]++
}

// 结构体嵌入和匿名成员
type Circle3 struct {
	X, Y, Radius int
}

type Circle2 struct {
	Center 	Point
	Radius int
}

type Wheel2 struct {
	Circle	Circle 
	Spokes 	int
}

// 结构体的匿名成员
type Circle struct {
	Point
	Radius 		int
}

type Wheel struct {
	Circle
	Spokes		int
}

func embed(){
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		}, 
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)
}


func main()  {
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)
	id := dilbert.ID
	// 下面语句通过EmployeeByID返回的结构体指针更新了Employee结构体的成员
	EmployeeByID(id).Salary = 0
	arr := []int{112, 38, 3, 22, -5}
	fmt.Println(Sort(arr))

	fmt.Println(Scale(Point{1, 2}, 5))

	equalStruct()
	embed()
}