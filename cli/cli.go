package cli

import (
	"fmt"
	"os"
)

// positionFound check a position is defined positions
func positionFound(positions *[]int, position int) bool {
	for _, i := range *positions {
		if position == i {
			return true
		}
	}

	return false
}

func positionsOutOfBoundary(positions *[]int, max int) bool {
	for _, i := range *positions {
		if i > max {
			return true
		}
	}

	return false
}

// checkFatalError is an helper to print an error and exit
func checkFatalError(err error) {
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
