import (
	"image/color"
)

type Point struct { X, Y float64 }

type ColoredPoint struct {
	Point
	Color	color.RGBA
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p * ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

// 下面是一个小trick。这个例子展示了简单的cache，其使用两个包级别的变量来实现，一个mutex互斥量(§9.2)和它所操作的cache
var (
	mu sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// 下面这个版本在功能上是一致的，但将两个包级别的变量放在了cache这个struct一组内
var cache = struct{
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string)
}

func Lookup2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}