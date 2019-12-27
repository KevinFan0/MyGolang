package geometry

import (
	"math"
)

type Point struct { X, Y float64 }

// traditional function
func Distance(p, q Point) float64  {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 方法调用
p := Point{1, 2}
q := Point{4, 6}
fmt.Println(Distance(p, q))
fmt.Println(p.Distance(q))

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}