package tokenizer

import (
	"fmt"
	"github.com/antham/goller/parser"
)

type Token struct {
	Value string
}

// Tokenizer handle string tokenization
type Tokenizer struct {
	parse         *parser.Parser
	maxTokensSize int
}

// NewTokenizer instantiate sequence objects
func NewTokenizer(parse *parser.Parser) *Tokenizer {
	return &Tokenizer{
		parse: parse,
	}
}

// Tokenize split a line to tokens
func (t *Tokenizer) Tokenize(line string) ([]Token, error) {
	var tokens []Token

	for _, data := range (*t.parse)(line) {
		tokens = append(tokens, Token{Value: data})
	}

	if (*t).maxTokensSize == 0 {
		(*t).maxTokensSize = len(tokens)
	}

	if len(tokens) != (*t).maxTokensSize {
		err := fmt.Errorf("Wrong parsing strategy (based on first line tokenization), got %d tokens instead of %d\nLine : %s\n", len(tokens), (*t).maxTokensSize, line)

		return []Token{}, err
	}

	return tokens, nil
}
