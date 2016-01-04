package transformer

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

	switch ch {
	case eof:
		return EOF, ""
	case '|':
		return PIPE, string(ch)
	case '(':
		return OPAREN, string(ch)
	case ')':
		return CPAREN, string(ch)
	case ':':
		return COLON, string(ch)
	case '"':
		return DQUOTE, string(ch)
	case ',':
		return COMMA, string(ch)
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
		if c == '"' || c == '|' || c == ')' || c == '(' || c == ',' || c == ':' || c == eof {
			s.unread()
			break
		} else if isLetter(c) {
			hasLetter = true
		} else if isNumber(c) {
			hasNumber = true
		} else {
			hasSpecialChar = true
		}

		data += string(c)
	}

	if hasNumber && !hasLetter && !hasSpecialChar {
		return NUMBER, data
	} else if (hasLetter || hasNumber) && !hasSpecialChar {
		return ALNUM, data
	}

	return STRING, data
}

// isLetter check if character is letter character
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isNumber check if character is number character
func isNumber(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
