package transformer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendUnexistingFunction(t *testing.T) {
	transformers := NewTransformers()

	err := transformers.Append(1, "whatever", []string{})

	if err == nil || err.Error() != `"whatever" doesn't exist or number of arguments "0" is wrong` {
		t.Errorf("Must throws an error")
	}
}

func TestLow(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "low", []string{}))

	result := transformers.Apply(1, "A RANDOM TEST")
	expected := "a random test"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestUpp(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "upp", []string{}))

	result := transformers.Apply(1, "a random test")
	expected := "A RANDOM TEST"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestTrim(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "trim", []string{"1"}))

	result := transformers.Apply(1, "11test11")
	expected := "test"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestTrimLeft(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "triml", []string{"1"}))

	result := transformers.Apply(1, "11test11")
	expected := "test11"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestTrimRight(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "trimr", []string{"1"}))

	result := transformers.Apply(1, "11test11")
	expected := "11test"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestReplace(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "repl", []string{"test", "hello world !"}))

	result := transformers.Apply(1, "test")
	expected := "hello world !"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestRightConcat(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "catr", []string{" world"}))

	result := transformers.Apply(1, "hello")
	expected := "hello world"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestLeftConcat(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "catl", []string{"hello"}))

	result := transformers.Apply(1, " world")
	expected := "hello world"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestStringLength(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "len", []string{}))

	result := transformers.Apply(1, "hello world")
	expected := "11"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestMatchWithBadRegexp(t *testing.T) {
	transformers := NewTransformers()

	err := transformers.Append(1, "match", []string{"(."})

	if err == nil || err.Error() != "An error occurred when parsing regexp : \"error parsing regexp: missing closing ): `(.`\"" {
		t.Errorf("Must throws an error")
	}
}

func TestMatchSuccessful(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "match", []string{"hello.*"}))

	result := transformers.Apply(1, "hello world")
	expected := "true"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestMatchFailed(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "match", []string{"a"}))

	result := transformers.Apply(1, "hello world")
	expected := "false"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestAdd(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "add", []string{"1"}))

	result := transformers.Apply(1, "2")
	expected := "3"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestAddWithNonIntegerArg(t *testing.T) {
	transformers := NewTransformers()

	err := transformers.Append(1, "add", []string{"a"})

	if err == nil || err.Error() != "Argument must be an integer, \"a\" given" {
		t.Errorf("Must throws an error")
	}
}

func TestSubstract(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "sub", []string{"1"}))

	result := transformers.Apply(1, "2")
	expected := "1"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestSubWithNonInteger(t *testing.T) {
	transformers := NewTransformers()

	err := transformers.Append(1, "sub", []string{"a"})

	if err == nil || err.Error() != "Argument must be an integer, \"a\" given" {
		t.Errorf("Must throws an error")
	}
}

func TestDeleteNumberOfCharactersAtTheRightSide(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "delr", []string{"8"}))

	result := transformers.Apply(1, "Hello world !")
	expected := "Hello"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestDeleteNumberOfCharactersAtTheRightSideWithNonInteger(t *testing.T) {
	transformers := NewTransformers()

	err := transformers.Append(1, "delr", []string{"a"})

	if err == nil || err.Error() != "Argument must be an integer, \"a\" given" {
		t.Errorf("Must throws an error")
	}
}

func TestDeleteNumberOfCharactersAtTheLeftSide(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "dell", []string{"6"}))

	result := transformers.Apply(1, "Hello world !")
	expected := "world !"

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestDeleteNumberOfCharactersAtTheLeftSideWithNonInteger(t *testing.T) {
	transformers := NewTransformers()

	err := transformers.Append(1, "dell", []string{"a"})

	if err == nil || err.Error() != "Argument must be an integer, \"a\" given" {
		t.Errorf("Must throws an error")
	}
}

func TestDeleteBiggerNumberOfCharactersTheLeftSide(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "dell", []string{"100"}))

	result := transformers.Apply(1, "Hello world !")
	expected := ""

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestDeleteBiggerNumberOfCharactersTheRightSide(t *testing.T) {
	transformers := NewTransformers()

	assert.NoError(t, transformers.Append(1, "delr", []string{"100"}))

	result := transformers.Apply(1, "Hello world !")
	expected := ""

	if result != expected {
		t.Errorf("%s got %s", expected, result)
	}
}
