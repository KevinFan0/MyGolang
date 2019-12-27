p := Point{1, 2}
q := Point{4, 6}

distanceFromP := p.Distance
fmt.Println(distanceFromP(q))
var origin Point
fmt.Println(distanceFromP(origin))

scaleP := p.ScaleBy
scaleP(2)
scaleP(3)
scaleP(10)

type Rocket struct {}
func (r *Rocket) Launch() {
	
}

r := new(Rocket)
time.AfterFunc(10 * time.Second, func() { r.Launch() })

// 也可直接用方法值传入AfterFunc更为简短
time.AfterFunc(10 * time.Second, r.Launch)

// 方法表达式

distance := Point.Distance
fmt.Println(distance(p, q))
fmt.Printf("%T\n", distance)

scale := (*Point).ScaleBy
scale(&p, 2)
fmt.Println(p)
fmt.Printf("%T\n", scale)


type Point struct { X, Y float64 }

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path	[]Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	}else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}