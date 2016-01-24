package reader

import (
	"strings"
	"testing"
)

func TestReadStdin(t *testing.T) {
	input := strings.NewReader("test1\ntest2\ntest3")

	entries := []string{
		"test1",
		"test2",
		"test3",
	}

	r := Reader{
		Input: input,
	}

	r.Read(func(line string) {
		expected := entries[0]

		if expected != line {
			t.Errorf("Line must be %s, got %s", expected, line)
		}

		if len(entries) > 0 {
			entries = entries[1:]
		}
	})
}
