package tokenizer

import (
	"github.com/antham/goller/parser"
	"testing"
)

func TestTokenizeLine(t *testing.T) {
	tok := NewTokenizer(nil)

	tokens := tok.Tokenize("80.154.42.54 - - [23/Aug/2010:15:25:35 +0000] \"GET /phpmy-admin/scripts/setup.php HTTP/1.1\" 404 347 \"-\" \"ZmEu\"")

	if tokens == nil {
		t.Error("tokens can't be nil")
	}

	if len(tokens) != 21 {
		t.Error("Expected length is 21 got", len(tokens))
	}

	if tokens[0].Value != "80.154.42.54" {
		t.Error("We should retrieve an IP 80.154.42.54 at token 0, got", tokens[0].Value)
	}

	if tokens[19].Value != "ZmEu" {
		t.Error("We should retrieve a string ZmEu at token 19, got", tokens[19].Value)
	}
}

func TestTokenizeLineWithAParser(t *testing.T) {
	p := new(parser.Parser)
	p.Create("whi", []string{})

	tok := NewTokenizer(*p)

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
