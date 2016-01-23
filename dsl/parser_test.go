package dsl

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParsePositionAndFunctions(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1.test2.test3(\"@[{}_ |,(hello world),| _{}]@\",\"\\\\whatever\\\").test4.test5(\"hello\").test6.test7(\"hello\",\"world\",\"!\")"))
	stmt, err := parser.ParsePositionAndFunctions()

	expected := &PositionStatement{
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

func TestParsePositionAndFunctionsUnValidPosition(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("a:test1.test2.test3"))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \"a\", expected a number" {
		t.Error("Must throw an error if no position is found")
	}
}

func TestParsePositionAndFunctionsUnvalidColon(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8|test1.test2.test3"))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \"|\", expected a colon" {
		t.Error("Must throw an error if no colon is found")
	}
}

func TestParsePositionAndFunctionsUnvalidFunction(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1@.test2.test3"))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \"test1@\", function must have letter and number only" {
		t.Error("Must throw an error if function as an invalid format")
	}
}

func TestParsePositionAndFunctionsNoDot(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1.test2.test3(\"arg1\"),"))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \",\", chainer must be a dot" {
		t.Error("Must throw an error if no colon is found")
	}
}

func TestParseFunctionPositionAndFunctionsWithNoDoubleQuoteToStartArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1.test2(hello).test3"))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \"hello\", arg must start with a quote" {
		t.Error("Must throw an error if function args doesn't start with a double quote")
	}
}

func TestParseFunctionPositionAndFunctionsWithNoDoubleQuoteToEndArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1.test2(\"hello"))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \"hello\", arg must end with a quote" {
		t.Error("Must throw an error if function args doesn't end with a double quote")
	}
}

func TestParseFunctionPositionAndFunctionsWithNoFinalParenToEndArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1.test2(\"hello\""))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \"\", must be a comma or close paren" {
		t.Error("Must throw an error if function arg doesn't end with paren")
	}
}

func TestParseFunctionPositionAndFunctionsWithNoCommaAfterArg(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1.test2(\"hello\".\"world\")"))
	stmt, err := parser.ParsePositionAndFunctions()

	if stmt != nil || err == nil || err.Error() != "found \".\", must be a comma or close paren" {
		t.Error("Must throw an error if function args are not separated with comma")
	}
}

func TestParseFunctionWithoutArgs(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("test1"))
	stmt, err := parser.ParseFunction()

	expected := &FunctionStatement{
		Name: "test1",
		Args: []string{},
	}

	if err != nil || reflect.DeepEqual(stmt, expected) != true {
		t.Errorf("Struct not equals expected %v got %v", expected, stmt)
	}
}

func TestParseFunctionWithArgs(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("test1(\"hello\",\"world\")"))
	stmt, err := parser.ParseFunction()

	expected := &FunctionStatement{
		Name: "test1",
		Args: []string{"hello", "world"},
	}

	if err != nil || reflect.DeepEqual(stmt, expected) != true {
		t.Errorf("Struct not equals expected %v got %v", expected, stmt)
	}
}

func TestParseFunctionWithExtraCharacters(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("test1(\"hello\",\"world\")|"))
	_, err := parser.ParseFunction()

	if err == nil || err.Error() != "found \"|\", only one function can be defined" {
		t.Error("Must throw an error if characters remain after single function definition")
	}
}

func TestParseFunctionUnvalidFunction(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("test1(\"hello\","))
	_, err := parser.ParseFunction()

	if err == nil || err.Error() != "found \"\", arg must start with a quote" {
		t.Error("Must throw an error if function is not correct")
	}
}

func TestParsePositionsAndFunctions(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1(\"1\"),9:test2(\"2\"),10:test3"))
	stmt, err := parser.ParsePositionsAndFunctions()

	expected := &[]PositionStatement{
		{
			Position: 8,
			Functions: []FunctionStatement{
				{
					Name: "test1",
					Args: []string{"1"},
				},
			},
		},
		{
			Position: 9,
			Functions: []FunctionStatement{
				{
					Name: "test2",
					Args: []string{"2"},
				},
			},
		},
		{
			Position: 10,
			Functions: []FunctionStatement{
				{
					Name: "test3",
					Args: []string{},
				},
			},
		},
	}

	if err != nil || reflect.DeepEqual(stmt, expected) != true {
		t.Errorf("Struct not equals expected %v got %v", expected, stmt)
	}
}

func TestParsePositionsAndFunctionsWithUnvalidPosition(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("test1(\"hello\")"))
	_, err := parser.ParsePositionsAndFunctions()

	if err == nil || err.Error() != "found \"test1\", expected a number" {
		t.Error("Must throw an error if position is not correct")
	}
}

func TestParsePositionsAndFunctionsWithUnvalidFunction(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test*(\"hello\")"))
	_, err := parser.ParsePositionsAndFunctions()

	if err == nil || err.Error() != "found \"test*\", function must have letter and number only" {
		t.Error("Must throw an error if function is not correct")
	}
}

func TestParsePositionsAndFunctionsEndingWithUnauthorizedCharacter(t *testing.T) {
	parser := NewParser(bytes.NewBufferString("8:test1|"))
	_, err := parser.ParsePositionsAndFunctions()

	if err == nil || err.Error() != "found \"|\", must be a comma" {
		t.Error("Must throw an error if position is not correct")
	}
}
