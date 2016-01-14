package cli

import (
	"testing"
)

func TestTransformersSetValidArgument(t *testing.T) {
	trans := &Transformers{}

	if trans.Set("8:low") != nil {
		t.Error("Must return no error")
	}
}

func TestTransformersSetUnValidArgument(t *testing.T) {
	trans := &Transformers{}

	if trans.Set("8:low(test)").Error() != "found \"test\", arg must start with a quote" {
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
