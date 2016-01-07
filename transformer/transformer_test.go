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

func TestTrimRight(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("trimr", []string{"1"})

	result := transformers.Apply("11test11")
	expected := "11test"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestReplace(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("repl", []string{"test", "hello world !"})

	result := transformers.Apply("test")
	expected := "hello world !"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestConcat(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("cat", []string{" world"})

	result := transformers.Apply("hello")
	expected := "hello world"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestStringLength(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("len", []string{})

	result := transformers.Apply("hello world")
	expected := "11"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}
