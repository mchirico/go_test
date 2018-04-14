package shapes

import (
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func CreateRect(width float64, height float64) rect {

	return rect{width: width, height: height}
}

func CreateCircle(radius float64) circle {

	return circle{radius: radius}
}

func Measure(g Geometry) (float64, float64) {

	return g.area(), g.perim()

}
