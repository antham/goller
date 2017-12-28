package dsl

import (
	"bufio"
	"io"
	"strconv"
)

var eof = rune(0)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan extract token from characters
func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	chs := map[rune]Token{
		eof: EOF,
		'|': PIPE,
		'(': OPAREN,
		')': CPAREN,
		':': COLON,
		'"': DQUOTE,
		',': COMMA,
		'.': DOT,
	}

	if v, ok := chs[ch]; ok {
		return v, string(ch)
	}

	if ch == '\\' {
		if s.read() == '"' {
			return EDQUOTE, string('"')
		}
		s.unread()
	}

	if strconv.IsPrint(ch) {
		s.unread()
		return s.scanString()
	}

	return ILLEGAL, string(ch)
}

// scanString return token according to string type
func (s *Scanner) scanString() (token Token, data string) {
	hasLetter := false
	hasNumber := false
	hasSpecialChar := false

	for c := s.read(); ; c = s.read() {
		if isMarker(c) {
			s.unread()
			break
		}

		switch true {
		case isLetter(c):
			hasLetter = true
		case isNumber(c):
			hasNumber = true
		default:
			hasSpecialChar = true
		}

		data += string(c)
	}

	if isStrNumber(hasNumber, hasLetter, hasSpecialChar) {
		return NUMBER, data
	}

	if isStrAlNum(hasNumber, hasLetter, hasSpecialChar) {
		return ALNUM, data
	}

	return STRING, data
}

// isStrNumber checks if a string is compound only with numbers
func isStrNumber(hasNumber, hasLetter, hasSpecialChar bool) bool {
	return hasNumber && !hasLetter && !hasSpecialChar
}

// isStrAlNum checks if a string is compound of numbers and letters
func isStrAlNum(hasNumber, hasLetter, hasSpecialChar bool) bool {
	return (hasLetter || hasNumber) && !hasSpecialChar
}

// isMarker checks rune is one of the defined runes
func isMarker(ch rune) bool {
	return (ch == '"' || ch == '|' || ch == ')' || ch == '(' || ch == ',' || ch == ':' || ch == '.' || ch == eof)
}

// isLetter checks if character is letter character
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isNumber checks if character is number character
func isNumber(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
