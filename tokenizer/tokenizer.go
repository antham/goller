package tokenizer

import (
	"github.com/antham/goller/parser"
	"github.com/trustpath/sequence"
)

// Tokenizer handle string tokenization
type Tokenizer struct {
	scanner *sequence.Scanner
	parser  parser.Parser
}

// Init instantiate sequence objects
func NewTokenizer(pars parser.Parser) *Tokenizer {
	return &Tokenizer{
		scanner: sequence.NewScanner(),
		parser:  pars,
	}
}

// Tokenize split a line to tokens
func (t *Tokenizer) Tokenize(line string) []sequence.Token {
	var tokens []sequence.Token

	if t.parser != nil {
		for _, data := range t.parser.Parse(line) {
			tokens = append(tokens, sequence.Token{Value: data})
		}

		return tokens
	}

	tokens, err := t.scanner.Scan(line)

	if err == nil {
		return tokens
	}

	return nil
}
