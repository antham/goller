package dsl

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseValidTransformer(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1|test2|test3(\"@[{}_ |,(hello world),| _{}]@\",\"\\\\whatever\\\")|test4|test5(\"hello\")|test6|test7(\"hello\",\"world\",\"!\")"))
	stmt, err := parser.Parse()
	t.Log("8:test1|test2|test3(\"@[{}_ |,(hello world),| _{}]@\",\"\\whatever\\\")|test4|test5(\"hello\")|test6|test7(\"hello\",\"world\",\"!\")")
	expected := &Statement{
		Position: 8,
		Functions: []FunctionStatement{
			{
				Name: "test1",
				Args: []string{},
			},
			{
				Name: "test2",
				Args: []string{},
			},
			{
				Name: "test3",
				Args: []string{
					"@[{}_ |,(hello world),| _{}]@",
					"\\whatever\\",
				},
			},
			{
				Name: "test4",
				Args: []string{},
			},
			{
				Name: "test5",
				Args: []string{
					"hello",
				},
			},
			{
				Name: "test6",
				Args: []string{},
			},
			{
				Name: "test7",
				Args: []string{
					"hello",
					"world",
					"!",
				},
			},
		},
	}

	if err != nil || reflect.DeepEqual(stmt, expected) != true {
		t.Errorf("Struct not equals expected %v got %v", expected, stmt)
	}
}

func TestParseUnValidPosition(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("a:test1|test2|test3"))
	stmt, err := parser.Parse()

	if stmt != nil || err == nil || err.Error() != "found \"a\", expected a number" {
		t.Error("Must throw an error if no position is found")
	}
}

func TestParseUnvalidColon(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8|test1|test2|test3"))
	stmt, err := parser.Parse()

	if stmt != nil || err == nil || err.Error() != "found \"|\", expected a colon" {
		t.Error("Must throw an error if no colon is found")
	}
}

func TestParseUnvalidFunction(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1@|test2|test3"))
	stmt, err := parser.Parse()

	if stmt != nil || err == nil || err.Error() != "found \"test1@\", function must have letter and number only" {
		t.Error("Must throw an error if function as an invalid format")
	}
}

func TestParseFunctionWithNoDoubleQuoteToStartArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1|test2(hello)|test3"))
	stmt, err := parser.Parse()

	if stmt != nil || err == nil || err.Error() != "found \"hello\", arg must start with a quote" {
		t.Error("Must throw an error if function args doesn't start with a double quote")
	}
}

func TestParseFunctionWithNoDoubleQuoteToEndArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1|test2(\"hello"))
	stmt, err := parser.Parse()

	if stmt != nil || err == nil || err.Error() != "found \"hello\", arg must end with a quote" {
		t.Error("Must throw an error if function args doesn't end with a double quote")
	}
}

func TestParseFunctionWithNoFinalParenToEndArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1|test2(\"hello\""))
	stmt, err := parser.Parse()

	if stmt != nil || err == nil || err.Error() != "found \"\", must be a comma or close paren" {
		t.Error("Must throw an error if function arg doesn't end with paren")
	}
}

func TestParseFunctionWithNoCommaAfterArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1|test2(\"hello\"|\"world\")"))
	stmt, err := parser.Parse()

	if stmt != nil || err == nil || err.Error() != "found \"|\", must be a comma or close paren" {
		t.Error("Must throw an error if function args are not separated with comma")
	}
}
