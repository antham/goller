package cli

import (
	"reflect"
	"testing"
)

func TestParserWrapper(t *testing.T) {
	result := ParserWrapper(MockSettings{})

	got := reflect.TypeOf(result).String()
	expected := "*cli.Parser"

	if got != expected {
		t.Errorf("Must return %s, got %s", expected, got)
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
