package cli

import (
	"reflect"
	"testing"
)

func TestTransformersWrapper(t *testing.T) {
	result := TransformersWrapper(MockSettings{})

	got := reflect.TypeOf(result).String()
	expected := "*cli.Transformers"

	if got != expected {
		t.Errorf("Must return %s, got %s", expected, got)
	}
}

func TestTransformersSetValidArgument(t *testing.T) {
	trans := &Transformers{}

	if trans.Set("8:low") != nil {
		t.Error("Must return no error")
	}

	ts := trans.Get()

	if len(*ts) != 1 {
		t.Errorf("Must have a length of 1, got %d", len(*ts))
	}

	got := (*ts)[8][0]("TEST")
	expected := "test"

	if got != expected {
		t.Errorf("Must return %s, got %s", expected, got)
	}
}

func TestTransformersSetUnValidFunction(t *testing.T) {
	trans := &Transformers{}

	if trans.Set("8:low(test)").Error() != "found \"test\", arg must start with a quote" {
		t.Error("Must return an error")
	}
}

func TestTransformersSetUnValidArgument(t *testing.T) {
	trans := &Transformers{}

	if trans.Set("8:add(\"a\")").Error() != "Argument must be an integer, \"a\" given" {
		t.Error("Must return an error")
	}
}

func TestTransformersSetUnexistingPosition(t *testing.T) {
	positions := []int{8}

	trans := &Transformers{}

	if trans.Set("9:low") != nil {
		t.Error("Must return no error")
	}

	if trans.ValidatePositions(&positions) != nil && trans.ValidatePositions(&positions).Error() != "Transformer is wrong : position 9 doesn't exist" {
		t.Error("Must return an error")
	}
}

func TestTransformersSetNoTransformers(t *testing.T) {
	positions := []int{8}

	trans := &Transformers{}

	if trans.ValidatePositions(&positions) != nil && trans.ValidatePositions(&positions).Error() != "Transformer is wrong : position 9 doesn't exist" {
		t.Error("Must return an error")
	}
}

func TestTransformersSetTransformerAtPosition0(t *testing.T) {
	trans := &Transformers{}

	if trans.Set("0:low").Error() != "You cannot add a transformer to position 0" {
		t.Error("Must return an error")
	}
}

func TestTransformersIsCumulative(t *testing.T) {
	trans := &Transformers{}

	if !trans.IsCumulative() {
		t.Error("Must return true")
	}
}

func TestTransformersString(t *testing.T) {
	trans := &Transformers{}

	if trans.String() != "" {
		t.Error("Must return an empty string")
	}
}
