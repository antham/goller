package cli

import (
	"fmt"
	"github.com/antham/goller/reader"
	"github.com/antham/goller/tokenizer"
	"strconv"
)

// tokenize is tied to tokenize command
type tokenize struct {
	tokenizer tokenizer.Tokenizer
	reader    reader.Reader
	tokens    []tokenizer.Token
}

// NewTokenize create an object related to tokenize command
func NewTokenize(args *tokenizeCommand) *tokenize {
	return &tokenize{
		tokenizer: *tokenizer.NewTokenizer(args.parser.Get()),
		reader:    reader.NewStdinReader(),
	}
}

// Consume tokenize every line from reader
func (p *tokenize) Tokenize() {
	p.reader.ReadFirstLine(func(line string) {
		p.tokens, _ = p.tokenizer.Tokenize(line)
	})
}

// Render tokens
func (p *tokenize) Render() {
	for i, token := range p.tokens {
		fmt.Println("position " + strconv.Itoa(i+1) + " => " + token.Value)
	}
}
