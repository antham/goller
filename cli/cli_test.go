package cli

import (
	"errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"testing"
)

type MockValue struct {
}

func (m MockValue) String() string {
	return ""
}

func (m MockValue) Set(string) error {
	return errors.New("")
}

type MockSettings struct {
}

func (m MockSettings) SetValue(value kingpin.Value) {
}

func TestPositionFoundWithUnexistingPosition(t *testing.T) {
	positions := []int{1, 2, 3}

	result := positionFound(&positions, 4)

	if result != false {
		t.Error("Position 4 must not exists")
	}
}
