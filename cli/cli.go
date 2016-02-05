package cli

import (
	"fmt"
	"os"
)

// positionsOutOfBoundary check if positions don't overflow boundaries
func positionsOutOfBoundary(positions *[]int, max int) bool {
	for _, i := range *positions {
		if i > max {
			return true
		}
	}

	return false
}

// triggerFatalError is an helper to print an error and exit
func triggerFatalError(err error) {
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
