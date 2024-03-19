package calculator_test

import (
	"testing"

	"github.com/kscott5/calculator"
)

func TestSetSize(t *testing.T) {
	triangle := calculator.Triangle{}
	triangle.SetSize(3)

	expect := 3
	actual := triangle.Size()

	if expect != triangle.Size() {
		t.Errorf("Triangle.Size: expected %d got %d", expect, actual)
	}
}

func TestPerimeter(t *testing.T) {
	triangle := calculator.Triangle{}
	triangle.SetSize(3)

	expect := 18
	actual := triangle.Perimeter()

	if expect != actual {
		t.Errorf("Triangle.Perimeter: expected %d got %d", expect, actual)
	}
}
