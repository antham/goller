package tokenizer

import (
	"testing"
)

func TestTokenizeLine(t *testing.T) {
	Init()

	tokens := Tokenize("80.154.42.54 - - [23/Aug/2010:15:25:35 +0000] \"GET /phpmy-admin/scripts/setup.php HTTP/1.1\" 404 347 \"-\" \"ZmEu\"")

	if tokens == nil {
		t.Error("tokens can't be nil")
	}

	if len(tokens) != 21 {
		t.Error("Expected length is 21 got", len(tokens))
	}

	if tokens[0].Value != "80.154.42.54" {
		t.Error("We should retrieve an IP 80.154.42.54 at token 0, got", tokens[0].Value)
	}

	if tokens[19].Value != "zmeu" {
		t.Error("We should retrieve a string zmeu at token 19, got", tokens[19].Value)
	}
}
