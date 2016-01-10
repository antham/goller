package parser

import (
	"reflect"
	"testing"
)

func TestSetValidArgument(t *testing.T) {
	parser := new(Parser)

	if parser.Set("whi") != nil {
		t.Error("Must return no error")
	}
}

func TestSetUnValidArgument(t *testing.T) {
	parser := new(Parser)
	err := parser.Set("whi(")

	if err == nil || err.Error() != "found \"\", arg must start with a quote" {
		t.Error("Must throw an error")
	}
}

func TestString(t *testing.T) {
	parser := new(Parser)

	if parser.String() != "" {
		t.Error("Must return an empty string")
	}
}

func TestParseWhitespace(t *testing.T) {
	parser := new(Parser)
	parser.Create("whi", []string{})

	result := parser.Parse("hello world\t, a    testing  sentence !")
	expected := []string{"hello", "world", ",", "a", "testing", "sentence", "!"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}
