package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// extractPositions split positions fields from string
func extractPositions(fields string, size int) ([]int, error) {
	var positions []int

	if fields != "" {
		positionDups := make(map[int]bool, 0)

		for _, value := range strings.Split(fields, ",") {
			if position, err := strconv.Atoi(value); err == nil {
				if _, ok := positionDups[position]; ok == true {
					return []int{}, fmt.Errorf("This element is duplicated : %d", position)
				}

				if position >= size+1 {
					return []int{}, fmt.Errorf("Position %d is greater or equal than maximum position %d", position, size+1)
				}

				positionDups[position] = true
				positions = append(positions, position)
			} else {
				return []int{}, fmt.Errorf("%s is not a number", value)
			}
		}
	} else {
		return []int{}, fmt.Errorf("At least 1 element is required")
	}

	return positions, nil
}

func checkFatalError(err error) {
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
