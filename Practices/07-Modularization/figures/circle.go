package figures

import "math"

type Circle struct {
	Radio float32
}

func (c *Circle) getArea() float32 {
	return math.Pi * (c.Radio * c.Radio)
}

func (c *Circle) getPerimeter() float32 {
	return 2 * math.Pi * c.Radio
}
