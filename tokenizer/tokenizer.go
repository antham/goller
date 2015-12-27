package tokenizer

import (
	"github.com/trustpath/sequence"
)

var parser *sequence.Parser
var analyzer *sequence.Analyzer
var scanner *sequence.Scanner

//Init instantiate sequence objects
func Init() {
	sequence.ReadConfig("")
	parser = sequence.NewParser()
	scanner = sequence.NewScanner()
}

//Tokenize split a line to tokens
func Tokenize(line string) []sequence.Token {
	seq, err := scanner.Scan(line)

	if err == nil {
		if _, err := parser.Parse(seq); err != nil {
			return seq
		}
	}

	return nil
}
