package tokenizer

import (
	"github.com/trustpath/sequence"
)

var scanner *sequence.Scanner

// Init instantiate sequence objects
func Init() {
	scanner = sequence.NewScanner()
}

// Tokenize split a line to tokens
func Tokenize(line string) []sequence.Token {
	seq, err := scanner.Scan(line)

	if err == nil {
		return seq
	}

	return nil
}
