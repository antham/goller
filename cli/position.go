package cli

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"strconv"
	"strings"
)

// Positions is an array of position
type Positions struct {
	positions []int
}

// Set is used to populate statement from string
func (p *Positions) Set(value string) error {
	positionDups := map[int]bool{}

	for _, field := range strings.Split(value, ",") {
		if position, err := strconv.Atoi(field); err == nil {
			if _, ok := positionDups[position]; ok {
				return fmt.Errorf("This element is duplicated : %d", position)
			}
			positionDups[position] = true
			p.positions = append(p.positions, position)
		} else {
			return fmt.Errorf("%s is not a number", field)
		}
	}

	return nil
}

// Get positions
func (p *Positions) Get() *[]int {
	return &p.positions
}

// String
func (p *Positions) String() string {
	return ""
}

// PositionsWrapper is used to transform argument from command line
func PositionsWrapper(s kingpin.Settings) (target *Positions) {
	target = &Positions{}
	s.SetValue(target)
	return
}
