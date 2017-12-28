package parser

import (
	"log"
	"reflect"
	"testing"
)

func TestParseWhitespace(t *testing.T) {
	p, err := NewParser("whi", []string{})

	if err != nil {
		log.Fatal(err)
	}

	result := (*p)("hello world\t, a    testing  sentence !")
	expected := []string{"hello", "world", ",", "a", "testing", "sentence", "!"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestParseRegexpWithPatternMatchingWholeExpression(t *testing.T) {
	p, err := NewParser("reg", []string{"(h.{4}) (w.{4}), (a) (testing) (sentence) !"})

	if err != nil {
		log.Fatal(err)
	}

	result := (*p)("hello world, a testing sentence !")
	expected := []string{"hello", "world", "a", "testing", "sentence"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestParseRegexpWithPatternNotMatchingExpression(t *testing.T) {
	p, err := NewParser("reg", []string{"(h.{4}) (w.{4}) (testing) (sentence)"})

	if err != nil {
		log.Fatal(err)
	}

	result := (*p)("hello world, a testing sentence !")
	expected := []string{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestParseSplit(t *testing.T) {
	p, err := NewParser("spl", []string{"separator"})

	if err != nil {
		log.Fatal(err)
	}

	result := (*p)("helloseparatorworld,separatoraseparatortestingseparatorsentenceseparator!")
	expected := []string{"hello", "world,", "a", "testing", "sentence", "!"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestParseCommonLogFormat(t *testing.T) {
	p, err := NewParser("clf", []string{})

	if err != nil {
		log.Fatal(err)
	}

	result := (*p)(`127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326`)
	expected := []string{"127.0.0.1", "user-identifier", "frank", "10/Oct/2000:13:55:36 -0700", "GET /apache_pb.gif HTTP/1.0", "200", "2326"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}

func TestParseCommonLogFormatWithWrongFormat(t *testing.T) {
	p, err := NewParser("clf", []string{})

	if err != nil {
		log.Fatal(err)
	}

	result := (*p)(`127.0.0 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326`)
	expected := []string{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%s got %s", expected, result)
	}
}
