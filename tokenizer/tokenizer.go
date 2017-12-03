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
	tokens        []Token
}

// NewTokenizer instantiate sequence objects
func NewTokenizer(parse *parser.Parser) *Tokenizer {
	return &Tokenizer{
		parse: parse,
	}
}

// reset all data structures involved in tokenization
func (t *Tokenizer) reset() {
	t.tokens = t.tokens[:0]
}

// Tokenize split a line to tokens
func (t *Tokenizer) Tokenize(line []byte) error {
	t.reset()

	for _, data := range (*t.parse)(string(line[:])) {
		t.tokens = append(t.tokens, Token{Value: data})
	}

	size := len(t.tokens)

	if (*t).maxTokensSize == 0 {
		(*t).maxTokensSize = size
	}

	if size != (*t).maxTokensSize {
		err := fmt.Errorf("Wrong parsing strategy (based on first line tokenization), got %d tokens instead of %d\nLine : %s\n", size, (*t).maxTokensSize, line)

		return err
	}

	return nil
}

// Get token accumulator
func (t *Tokenizer) Get() *[]Token {
	return &(t.tokens)
}

// Length return token count
func (t *Tokenizer) Length() int {
	return len(t.tokens)
}
