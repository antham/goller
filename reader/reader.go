package reader

import (
	"bufio"
	"io"
	"io/ioutil"
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
	accumulator := []byte{}

	datas, err := ioutil.ReadAll(r.Input)

	if err != nil {
		return err
	}

	size := len(datas)

	for i, data := range datas {
		if data != '\n' {
			accumulator = append(accumulator, data)
		}

		if data == '\n' || i == size-1 {
			err := rowReader(string(accumulator[:]))

			if err != nil {
				return err
			}

			accumulator = []byte{}
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
