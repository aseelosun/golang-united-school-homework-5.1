package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	// implement me
	return Point{s.start.x + int(s.a), s.start.y + int(s.a)}
}

func (s *Square) Area() uint {
	d := math.Sqrt(math.Pow(float64(s.start.x-s.End().x), 2) + math.Pow(float64(s.start.y-s.End().y), 2))
	//math.Pow(float64(s.start.x-s.End().x), 2)
	//math.Pow(float64(s.start.y-s.End().y),2)
	area := math.Pow(d, 2)
	return uint(area)
}

func (s *Square) Perimeter() uint {
	area := s.Area()
	return area / 4
}
