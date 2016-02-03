package reader

import (
	"bufio"
	"io"
	"os"
)

// Reader
type Reader struct {
	Input io.Reader
}

// NewStdinReader create a reader tied to stdin
func NewStdinReader() Reader {
	return Reader{
		Input: os.Stdin,
	}
}

// Read split entries per line
func (r Reader) Read(rowReader func(line string)) {
	scanner := bufio.NewScanner(r.Input)
	for scanner.Scan() {
		rowReader(scanner.Text())
	}
}

// ReadFirstLine split entries per line
func (r Reader) ReadFirstLine(rowReader func(line string)) {
	scanner := bufio.NewScanner(r.Input)
	scanner.Scan()
	rowReader(scanner.Text())
}
