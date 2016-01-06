package transformer

import (
	"testing"
)

func TestSetValidArgument(t *testing.T) {
	trans := &TransformersMap{}

	if trans.Set("8:low") != nil {
		t.Error("Must return no error")
	}
}

func TestSetUnValidArgument(t *testing.T) {
	trans := &TransformersMap{}

	if trans.Set("8:low(test)").Error() != "found \"test\", arg must start with a quote" {
		t.Error("Must return an error")
	}
}

func TestIsCumulative(t *testing.T) {
	trans := &TransformersMap{}

	if trans.IsCumulative() != true {
		t.Error("Must return true")
	}
}

func TestString(t *testing.T) {
	trans := &TransformersMap{}

	if trans.String() != "" {
		t.Error("Must return an empty string")
	}
}
