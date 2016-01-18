package cli

import (
	"reflect"
	"testing"
)

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

func TestTransformersSetUnValidArgument(t *testing.T) {
	trans := &Transformers{}

	if trans.Set("8:low(test)").Error() != "found \"test\", arg must start with a quote" {
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

	if trans.IsCumulative() != true {
		t.Error("Must return true")
	}
}

func TestTransformersString(t *testing.T) {
	trans := &Transformers{}

	if trans.String() != "" {
		t.Error("Must return an empty string")
	}
}

func TestParserSetValidArgument(t *testing.T) {
	parser := new(Parser)

	if parser.Set("whi") != nil {
		t.Error("Must return no error")
	}

	p := parser.Get()

	got := (*p)("hello world !")
	expected := []string{"hello", "world", "!"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got %v, expected %v", got, expected)
	}
}

func TestParserSetUnValidArgument(t *testing.T) {
	parser := new(Parser)
	err := parser.Set("whi(")

	if err == nil || err.Error() != "found \"\", arg must start with a quote" {
		t.Error("Must throw an error")
	}
}

func TestParserString(t *testing.T) {
	parser := new(Parser)

	if parser.String() != "" {
		t.Error("Must return an empty string")
	}
}

func TestExtractPositionsFromString(t *testing.T) {
	positions, err := ExtractPositions("2,5,0,8,16,12", 17)

	if err != nil {
		t.Error("Got an error", err)
	}

	if len(positions) != 6 {
		t.Error("Expected slice length of 6, got", len(positions))
	}

	if positions[0] != 2 {
		t.Error("First element must be 2, got", positions[0])
	}

	if positions[5] != 12 {
		t.Error("First element must be 12, got", positions[5])
	}
}

func TestExtractPositionsMustReturnUniquePositions(t *testing.T) {
	_, err := ExtractPositions("2,5,5,8,16,12", 17)

	if err == nil || err.Error() != "This element is duplicated : 5" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromEmptyString(t *testing.T) {
	_, err := ExtractPositions("", 17)

	if err == nil || err.Error() != "At least 1 element is required" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromStringContainingSomethingDifferentThanNumber(t *testing.T) {
	_, err := ExtractPositions("1,2,a,b,c", 17)

	if err == nil || err.Error() != "a is not a number" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromStringContainingPositionOverLimit(t *testing.T) {
	_, err := ExtractPositions("1,2,5,12,8,9", 11)

	if err == nil || err.Error() != "Position 12 is greater or equal than maximum position 12" {
		t.Error("An error must occur", err)
	}

	_, err = ExtractPositions("1,2,5,13,8,9", 11)

	if err == nil || err.Error() != "Position 13 is greater or equal than maximum position 12" {
		t.Error("An error must occur", err)
	}
}
