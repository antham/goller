package parser

import (
	"reflect"
	"testing"
)

func TestParseWhitespace(t *testing.T) {
	p := NewParser("whi", []string{})

	result := (*p)("hello world\t, a    testing  sentence !")
	expected := []string{"hello", "world", ",", "a", "testing", "sentence", "!"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestParseRegexpWithPatternMatchingWholeExpression(t *testing.T) {
	p := NewParser("reg", []string{"(h.{4}) (w.{4}), (a) (testing) (sentence) !"})

	result := (*p)("hello world, a testing sentence !")
	expected := []string{"hello", "world", "a", "testing", "sentence"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestParseRegexpWithPatternNotMatchingExpression(t *testing.T) {
	p := NewParser("reg", []string{"(h.{4}) (w.{4}) (testing) (sentence)"})

	result := (*p)("hello world, a testing sentence !")
	expected := []string{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}
