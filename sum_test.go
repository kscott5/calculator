package calculator_test

import "github.com/kscott5/calculator"
import "testing"

func TestSum(t *testing.T) {
	expect := 3
	actual := calculator.Sum(1,2)

	if expect != actual {
		t.Errorf("Sum: expected %d got %d", expect, actual)
	}
}

func TestAppName(t *testing.T) {
	expect := "Calculator"
	actual := calculator.AppName

	if expect != actual {
		t.Errorf("Appname: expected %s got %s", expect, actual)
	}
}