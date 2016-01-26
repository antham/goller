package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"reflect"
	"testing"
)

func TestPositionsWrapper(t *testing.T) {
	result := PositionsWrapper(&kingpin.ArgClause{})

	got := reflect.TypeOf(result).String()
	expected := "*cli.Positions"

	if got != expected {
		t.Errorf("Must return %s, got %s", expected, got)
	}
}

func TestPositionsSetValidArgument(t *testing.T) {
	parser := new(Positions)

	if parser.Set("1,2,3,4,5") != nil {
		t.Error("Must return no error")
	}

	got := *parser.Get()
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got %v, expected %v", got, expected)
	}
}

func TestPositionsSetUnValidArgument(t *testing.T) {
	parser := new(Positions)
	err := parser.Set("1,a")

	if err == nil || err.Error() != "a is not a number" {
		t.Error("Must throw an error")
	}
}

func TestPositionsSetDuplicatedArgument(t *testing.T) {
	parser := new(Positions)
	err := parser.Set("1,1")

	if err == nil || err.Error() != "This element is duplicated : 1" {
		t.Error("Must throw an error")
	}
}

func TestPositionsString(t *testing.T) {
	parser := new(Positions)

	if parser.String() != "" {
		t.Error("Must return an empty string")
	}
}
