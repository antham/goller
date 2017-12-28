package reader

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
)

func TestNewRead(t *testing.T) {
	r := NewStdinReader()

	if r.Input != os.Stdin {
		t.Error("r.input is plugged on os.Stdin")
	}
}

func TestRead(t *testing.T) {
	input := strings.NewReader("test1\ntest2\ntest3")

	entries := [][]byte{
		[]byte("test1"),
		[]byte("test2"),
		[]byte("test3"),
	}

	r := Reader{
		Input: input,
	}

	r.Read(func(line []byte) error {
		expected := entries[0]

		if !bytes.Equal(expected, line) {
			t.Errorf("Line must be %s, got %s", string(expected), string(line))
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

	err := r.Read(func(line []byte) error {
		return errors.New("Error from inner function")
	})

	if err == nil || err.Error() != "Error from inner function" {
		t.Error("Read must return error from inner function")
	}
}

func TestReadFirstLine(t *testing.T) {
	input := strings.NewReader("test1\ntest2\ntest3")

	expected := []byte("test1")

	r := Reader{
		Input: input,
	}

	r.ReadFirstLine(func(line []byte) error {
		if !bytes.Equal(expected, line) {
			t.Errorf("Line must be %s, got %s", expected, line)
		}

		return nil
	})
}
