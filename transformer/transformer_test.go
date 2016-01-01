package transformer

import (
	"testing"
)

func TestAppend(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("low", []string{})
	transformers.Append("upp", []string{})

	result := transformers.Apply("A RANDOM TEST")
	expected := "a random test"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}
