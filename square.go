package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (p *Point) End(x int, y int) Point {
	// implement me
	return Point{x: x, y: y}
}

func (s *Square) Area() uint {
	// implement me
}

func (s *Square) Perimeter() uint {
	// implement me
}
