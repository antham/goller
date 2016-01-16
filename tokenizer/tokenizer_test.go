package tokenizer

import (
	"github.com/antham/goller/parser"
	"testing"
)

func TestTokenizeLineWithAParser(t *testing.T) {
	p := parser.NewParser("whi", []string{})

	tok := NewTokenizer(p)

	tokens := tok.Tokenize("[2016-01-08 20:16] [ALPM] transaction started")

	if tokens == nil {
		t.Error("tokens can't be nil")
	}

	if len(tokens) != 5 {
		t.Errorf("Expected length is %v got %v", 5, len(tokens))
	}

	if tokens[0].Value != "[2016-01-08" {
		t.Errorf("We should retrieve %v got %v", "[2016-01-08", tokens[0].Value)
	}

	if tokens[4].Value != "started" {
		t.Errorf("We should retrieve %v at token 4, got %v", "started", tokens[4].Value)
	}
}
