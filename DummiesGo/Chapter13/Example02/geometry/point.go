package geometry

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Length() float64 {
	return math.Sqrt(math.Pow(p.X, 2.0) + math.Pow(p.Y, 2.0))
}
