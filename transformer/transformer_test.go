package transformer

import (
	"testing"
)

func TestSetValidArgument(t *testing.T) {
	trans := &TransformersMap{}

	if trans.Set("8:low") != nil {
		t.Error("Must return no error")
	}
}

func TestSetUnValidArgument(t *testing.T) {
	trans := &TransformersMap{}

	if trans.Set("8:low(test)").Error() != "found \"test\", arg must start with a quote" {
		t.Error("Must return an error")
	}
}

func TestIsCumulative(t *testing.T) {
	trans := &TransformersMap{}

	if trans.IsCumulative() != true {
		t.Error("Must return true")
	}
}

func TestString(t *testing.T) {
	trans := &TransformersMap{}

	if trans.String() != "" {
		t.Error("Must return an empty string")
	}
}

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

func TestRightConcat(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("rcat", []string{" world"})

	result := transformers.Apply("hello")
	expected := "hello world"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestLeftConcat(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("lcat", []string{"hello"})

	result := transformers.Apply(" world")
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

func TestMatchSuccessful(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("match", []string{"hello.*"})

	result := transformers.Apply("hello world")
	expected := "true"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestMatchFailed(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("match", []string{"a"})

	result := transformers.Apply("hello world")
	expected := "false"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestAdd(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("add", []string{"1"})

	result := transformers.Apply("2")
	expected := "3"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestSubstract(t *testing.T) {
	transformers := &Transformers{}

	transformers.Append("sub", []string{"1"})

	result := transformers.Apply("2")
	expected := "1"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}
