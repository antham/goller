package cli

import (
	"os"
	"testing"
)

func TestRunGroup(t *testing.T) {
	os.Args = []string{"goller", "group", "whi", "0"}

	Run("0.0.1")
}
