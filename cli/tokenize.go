package cli

import (
	"fmt"
	"github.com/antham/goller/v2/reader"
	"github.com/antham/goller/v2/tokenizer"
	"strconv"
)

// Tokenize is tied to tokenize command
type Tokenize struct {
	tokenizer tokenizer.Tokenizer
	reader    reader.Reader
	tokens    []tokenizer.Token
}

// NewTokenize creates an object related to tokenize command
func NewTokenize(args *tokenizeCommand) *Tokenize {
	return &Tokenize{
		tokenizer: *tokenizer.NewTokenizer(args.parser.Get()),
		reader:    reader.NewStdinReader(),
	}
}

// Tokenize tokenizes every line from reader
func (p *Tokenize) Tokenize() error {
	return p.reader.ReadFirstLine(func(line []byte) error {
		_ = p.tokenizer.Tokenize(line)
		p.tokens = *(p.tokenizer.Get())

		return nil
	})
}

// Render tokens
func (p *Tokenize) Render() {
	for i, token := range p.tokens {
		fmt.Println("position " + strconv.Itoa(i+1) + " => " + token.Value)
	}
}
