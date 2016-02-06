package reader

import (
	"errors"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	input := strings.NewReader("test1\ntest2\ntest3")

	entries := []string{
		"test1",
		"test2",
		"test3",
	}

	r := Reader{
		Input: input,
	}

	r.Read(func(line string) error {
		expected := entries[0]

		if expected != line {
			t.Errorf("Line must be %s, got %s", expected, line)
		}

		if len(entries) > 0 {
			entries = entries[1:]
		}

		return nil
	})
}

func TestReadWithAnError(t *testing.T) {
	input := strings.NewReader("test")

	r := Reader{
		Input: input,
	}

	err := r.Read(func(line string) error {
		return errors.New("Error from inner function")
	})

	if err == nil || err.Error() != "Error from inner function" {
		t.Error("Read must return error from inner function")
	}
}

func TestReadFirstLine(t *testing.T) {
	input := strings.NewReader("test1\ntest2\ntest3")

	expected := "test1"

	r := Reader{
		Input: input,
	}

	r.ReadFirstLine(func(line string) error {
		if expected != line {
			t.Errorf("Line must be %s, got %s", expected, line)
		}

		return nil
	})
}
