package model

import (
	"math"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Returns distance to a point
func (p *Point) Distance(to Point) float64 {
	return math.Abs(p.X-to.X) + math.Abs(p.Y-to.Y)
}

// represents a Point with distance, which is used to sort the results
type PointWithDistance struct {
	Point
	distance float64 `json:"-"`
}
