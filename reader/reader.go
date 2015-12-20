package reader

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

var filenames []string

// Set init filenames to use
func Set(pattern string) {
	var err error

	filenames, err = filepath.Glob(pattern)

	if err != nil {
		log.Fatalf("An error occured when retrieving filenames from pattern : %s", pattern)
	}
}

func Read(rowReader func(line string)) {
	for _, filename := range filenames {
		pointer, err := os.Open(filename)
		defer pointer.Close()

		if err != nil {
			log.Fatalf("Can't open %s", filename)
		}

		scanner := bufio.NewScanner(pointer)

		for scanner.Scan() {
			rowReader(scanner.Text())
		}
	}
}
