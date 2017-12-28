package dsl

import (
	"bytes"
	"testing"
)

func TestIsLetter(t *testing.T) {
	if !isLetter('A') {
		t.Error("A not seen as letter character")
	}

	if !isLetter('Z') {
		t.Error("Z not seen as letter character")
	}

	if !isLetter('M') {
		t.Error("M not seen as letter character")
	}

	if !isLetter('a') {
		t.Error("a not seen as letter character")
	}

	if !isLetter('z') {
		t.Error("z not seen as letter character")
	}

	if !isLetter('m') {
		t.Error("m not seen as letter character")
	}

	if isLetter('4') {
		t.Error("4 seen as letter character")
	}
}

func TestIsNumber(t *testing.T) {
	if !isNumber('0') {
		t.Error("0 not seen as letter character")
	}

	if !isNumber('9') {
		t.Error("9 not seen as letter character")
	}

	if !isNumber('5') {
		t.Error("5 not seen as letter character")
	}

	if isNumber('a') {
		t.Error("a seen as number")
	}
}

func TestReadNumber(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("90"))
	token, data := scanner.Scan()

	if token != NUMBER || data != "90" {
		t.Error("Token must be a number got", data)
	}
}

func TestReadAlphaNum(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("0a1b3"))
	token, data := scanner.Scan()

	if token != ALNUM || data != "0a1b3" {
		t.Error("Token must be alphanum got", data)
	}
}

func TestReadString(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("0a1b3@"))
	token, data := scanner.Scan()

	if token != STRING || data != "0a1b3@" {
		t.Error("Token must be string got", data)
	}
}

func TestReadPipe(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("|"))
	token, data := scanner.Scan()

	if token != PIPE || data != "|" {
		t.Error("Token must be a pipe got", data)
	}
}

func TestReadOpenParens(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("("))
	token, data := scanner.Scan()

	if token != OPAREN || data != "(" {
		t.Error("Token must be an open parens got", data)
	}
}

func TestReadCloseParens(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString(")"))
	token, data := scanner.Scan()

	if token != CPAREN || data != ")" {
		t.Error("Token must be a close parens got", data)
	}
}

func TestReadColon(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString(":"))
	token, data := scanner.Scan()

	if token != COLON || data != ":" {
		t.Error("Token must be a colon got", data)
	}
}

func TestReadComma(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString(","))
	token, data := scanner.Scan()

	if token != COMMA || data != "," {
		t.Error("Token must be a comma got", data)
	}
}

func TestReadDoubleQuote(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("\""))
	token, data := scanner.Scan()

	if token != DQUOTE || data != "\"" {
		t.Error("Token must be a double quote got", data)
	}
}

func TestReadEscapedDoubleQuote(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("\\\""))
	token, data := scanner.Scan()

	if token != EDQUOTE || data[0] != '"' {
		t.Error("Token must be an escaped double quote got", data)
	}
}

func TestReadEOF(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString(""))
	token, data := scanner.Scan()

	if token != EOF || data != "" {
		t.Error("Token must mark end of file got", data)
	}
}

func TestReadIllegalCharacter(t *testing.T) {
	scanner := NewScanner(bytes.NewBufferString("\007"))
	token, data := scanner.Scan()

	if token != ILLEGAL || data == "" {
		t.Error("Token must mark as illegal got", data)
	}
}
