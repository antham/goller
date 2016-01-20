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

func TestExtractPositionsFromString(t *testing.T) {
	positions, err := ExtractPositions("2,5,0,8,16,12", 17)

	if err != nil {
		t.Error("Got an error", err)
	}

	if len(positions) != 6 {
		t.Error("Expected slice length of 6, got", len(positions))
	}

	if positions[0] != 2 {
		t.Error("First element must be 2, got", positions[0])
	}

	if positions[5] != 12 {
		t.Error("First element must be 12, got", positions[5])
	}
}

func TestExtractPositionsMustReturnUniquePositions(t *testing.T) {
	_, err := ExtractPositions("2,5,5,8,16,12", 17)

	if err == nil || err.Error() != "This element is duplicated : 5" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromEmptyString(t *testing.T) {
	_, err := ExtractPositions("", 17)

	if err == nil || err.Error() != "At least 1 element is required" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromStringContainingSomethingDifferentThanNumber(t *testing.T) {
	_, err := ExtractPositions("1,2,a,b,c", 17)

	if err == nil || err.Error() != "a is not a number" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromStringContainingPositionOverLimit(t *testing.T) {
	_, err := ExtractPositions("1,2,5,12,8,9", 11)

	if err == nil || err.Error() != "Position 12 is greater or equal than maximum position 12" {
		t.Error("An error must occur", err)
	}

	_, err = ExtractPositions("1,2,5,13,8,9", 11)

	if err == nil || err.Error() != "Position 13 is greater or equal than maximum position 12" {
		t.Error("An error must occur", err)
	}
}
