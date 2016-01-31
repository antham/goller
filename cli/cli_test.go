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

func TestPositionsOutOfBoundariesWithAnIntegerInferiorToOnePosition(t *testing.T) {
	positions := []int{4, 5, 1, 7, 0, 2}

	if !positionsOutOfBoundary(&positions, 6) {
		t.Error("Must return true")
	}
}

func TestPositionsOutOfBoundariesWithAnIntegerSuperiorToAllPositions(t *testing.T) {
	positions := []int{4, 5, 1, 7, 0, 2}

	if positionsOutOfBoundary(&positions, 8) {
		t.Error("Must return false")
	}
}
