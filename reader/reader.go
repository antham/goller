package reader

import (
	"bufio"
	"io"
	"os"
)

var input io.Reader

// init initialize variables
func init() {
	input = os.Stdin
}

// ReadStdin split stdin per line
func ReadStdin(rowReader func(line string)) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rowReader(scanner.Text())
	}
}
