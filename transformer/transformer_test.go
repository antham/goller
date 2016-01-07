package transformer

import (
	"testing"
)

func TestLow(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("low", []string{})

	result := transformers.Apply("A RANDOM TEST")
	expected := "a random test"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestUpp(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("upp", []string{})

	result := transformers.Apply("a random test")
	expected := "A RANDOM TEST"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}
