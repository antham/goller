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

func TestTrim(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("trim", []string{"1"})

	result := transformers.Apply("11test11")
	expected := "test"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestTrimLeft(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("triml", []string{"1"})

	result := transformers.Apply("11test11")
	expected := "test11"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}
