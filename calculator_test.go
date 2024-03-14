package calculator_test

// Documentation of package location, GOPATH or GOROOT expected instruction not working
import (
	"testing"

	"github.com/kscott5/calculator"
)

func TestSum(t *testing.T) {
	expect := 3
	actual := calculator.Sum(1, 2)

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
