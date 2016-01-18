package sorter

import (
	"github.com/antham/goller/agregator"
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
func (s *Sorters) Append(previousPosition int, currentPosition int, sorterName string, args []string) {
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
			previousPosition: previousPosition,
			currentPosition:  currentPosition,
			less:             fun,
		}

		*s = append(*s, sorter)
	}
}

// Sort agregators using provided sorters
func (s *Sorters) Sort(agregators *agregator.Agregators) {
	for _, sor := range *s {
		sor.SetAgregators(agregators)
		sort.Sort(sor)
	}
}

// Sorter represents a sorter applied to agregators
type Sorter struct {
	agregators       *agregator.Agregators
	previousPosition int
	currentPosition  int
	less             less
}

// SetAgregators populate agregators
func (s *Sorter) SetAgregators(agregators *agregator.Agregators) {
	s.agregators = agregators
}

func (s *Sorter) Len() int {
	return len(*(s.agregators))
}

func (s *Sorter) Less(i, j int) bool {
	return s.less(*(*s.agregators)[i].DatasOrdered[s.currentPosition], *(*s.agregators)[j].DatasOrdered[s.currentPosition])
}

func (s *Sorter) Swap(i, j int) {
	(*s.agregators)[i], (*s.agregators)[j] = (*s.agregators)[j], (*s.agregators)[i]
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
