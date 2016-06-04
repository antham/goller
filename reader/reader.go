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
func (r Reader) Read(rowReader func(line *[]byte) error) error {
	datas, err := ioutil.ReadAll(r.Input)

	if err != nil {
		return err
	}

	size := len(datas)

	var lineStartPos int
	var lineEndPos int

	for i, data := range datas {
		if i == size-1 {
			lineEndPos = i + 1
		} else {
			lineEndPos = i
		}

		if data == '\n' || i == size-1 {
			l := datas[lineStartPos:lineEndPos]

			err := rowReader(&l)

			if err != nil {
				return err
			}

			lineStartPos = lineEndPos + 1
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
