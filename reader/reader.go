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
func (r Reader) Read(rowReader func(line string) error) error {
	scanner := bufio.NewScanner(r.Input)
	for scanner.Scan() {
		err := rowReader(scanner.Text())

		if err != nil {
			return err
		}
	}

	return nil
}

// ReadFirstLine split entries per line
func (r Reader) ReadFirstLine(rowReader func(line string) error) error {
	scanner := bufio.NewScanner(r.Input)
	scanner.Scan()
	return rowReader(scanner.Text())
}
