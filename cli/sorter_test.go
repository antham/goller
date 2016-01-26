package cli

import (
	"reflect"
	"testing"
)

func TestSortersWrapper(t *testing.T) {
	positions := []int{}
	result := SortersWrapper(MockSettings{}, &positions)

	got := reflect.TypeOf(result).String()
	expected := "*cli.Sorters"

	if got != expected {
		t.Errorf("Must return %s, got %s", expected, got)
	}
}

func TestSortersSetValidArgument(t *testing.T) {
	positions := []int{8}

	sorter := &Sorters{
		positions: &positions,
	}

	if sorter.Set("8:str") != nil {
		t.Error("Must return no error")
	}

	st := sorter.Get()

	if len(*st) != 1 {
		t.Errorf("Must have a length of 1, got %d", len(*st))
	}
}

func TestSortersSetUnValidArgument(t *testing.T) {
	sorter := &Sorters{}

	if sorter.Set("8:str(test)").Error() != "found \"test\", arg must start with a quote" {
		t.Error("Must return an error")
	}
}

func TestSortersSetUnexistingPosition(t *testing.T) {
	positions := []int{8}
	sorter := &Sorters{
		positions: &positions,
	}

	if sorter.Set("9:str").Error() != "Sort is wrong : position 9 doesn't exist" {
		t.Error("Must return an error")
	}
}

func TestSortersString(t *testing.T) {
	sorter := &Sorters{}

	if sorter.String() != "" {
		t.Error("Must return an empty string")
	}
}
