package sorter

import (
	"github.com/antham/goller/v2/aggregator"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Sorters list all sorters to apply to positions
type Sorters []*Sorter

// NewSorters create sorter list
func NewSorters() *Sorters {
	return &Sorters{}
}

// Append sort operation
func (s *Sorters) Append(position int, sorterName string, args []string) {
	var fun less

	switch len(args) {
	case 0:
		switch sorterName {
		case "int":
			fun = integer
		case "strl":
			fun = strLength
		case "str":
			fun = str
		}
	}

	if fun != nil {
		sorter := &Sorter{
			position: position,
			less:     fun,
		}

		*s = append(*s, sorter)
	}
}

// Sort aggregators using provided sorters
func (s *Sorters) Sort(aggregators *aggregator.Aggregators) {
	for i := len(*s) - 1; i >= 0; i-- {
		(*s)[i].SetAggregators(aggregators)
		sort.Stable((*s)[i])
	}
}

// Sorter represents a sorter applied to aggregators
type Sorter struct {
	aggregators *aggregator.Aggregators
	position    int
	less        less
}

// SetAggregators populate aggregators
func (s *Sorter) SetAggregators(aggregators *aggregator.Aggregators) {
	s.aggregators = aggregators
}

func (s *Sorter) Len() int {
	return len(*(s.aggregators))
}

func (s *Sorter) Less(i, j int) bool {
	return s.less(*(*s.aggregators)[i].DatasOrdered[s.position], *(*s.aggregators)[j].DatasOrdered[s.position])
}

func (s *Sorter) Swap(i, j int) {
	(*s.aggregators)[i], (*s.aggregators)[j] = (*s.aggregators)[j], (*s.aggregators)[i]
}

// HasPosition check if position exists
func (s *Sorter) HasPosition(position int) bool {
	return s.position == position
}

// GetPosition return element position
func (s *Sorter) GetPosition() int {
	return s.position
}

// less represents a function used in sorting
type less func(leftValue, rightValue string) bool

// integer sort integer values
func integer(leftValue, rightValue string) bool {
	leftValueInt, err := strconv.Atoi(leftValue)

	if err != nil {
		log.Fatalf("%s is not an integer", leftValue)
	}

	rightValueInt, err := strconv.Atoi(rightValue)

	if err != nil {
		log.Fatalf("%s is not an integer", rightValue)
	}

	return leftValueInt < rightValueInt
}

// str sort string
func str(leftValue, rightValue string) bool {
	return strings.Compare(leftValue, rightValue) == -1
}

// stringLength sort using string length
func strLength(leftValue, rightValue string) bool {
	return len(leftValue) < len(rightValue)
}
