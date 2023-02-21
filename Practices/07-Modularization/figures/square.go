package figures

type Square struct {
	Side float32
}

func (s *Square) getArea() float32 {
	return s.Side * 2
}

func (s *Square) getPerimeter() float32 {
	return s.Side * 4
}
