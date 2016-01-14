package tokenizer

import (
	"github.com/antham/goller/parser"
	"github.com/trustpath/sequence"
)

// Tokenizer handle string tokenization
type Tokenizer struct {
	scanner *sequence.Scanner
	parse   *parser.Parser
}

// NewTokenizer instantiate sequence objects
func NewTokenizer(parse *parser.Parser) *Tokenizer {
	return &Tokenizer{
		scanner: sequence.NewScanner(),
		parse:   parse,
	}
}

// Tokenize split a line to tokens
func (t *Tokenizer) Tokenize(line string) []sequence.Token {
	var tokens []sequence.Token

	if t.parse != nil {
		for _, data := range (*t.parse)(line) {
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
