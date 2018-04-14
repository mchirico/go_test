package shapes

import (
	"testing"
)

func TestShapes(t *testing.T) {

	g := []Geometry{}

	type result struct {
		a int
		p int
	}
	expectedValues := []result{}
	expectedValues = append(expectedValues, result{12, 14})
	expectedValues = append(expectedValues, result{78, 31})

	g = append(g, CreateRect(3, 4))
	g = append(g, CreateCircle(5))

	for i, shape := range g {
		area, perm := Measure(shape)

		if expectedValues[i].a != int(area) && expectedValues[i].p != int(perm) {
			t.Errorf("Failed ")
		}

	}

}
