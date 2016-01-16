package agregator

import (
	"crypto/sha1"
	"fmt"
	"github.com/antham/goller/transformer"
	"github.com/trustpath/sequence"
	"strconv"
	"strings"
)

// Agregator represents a unique log line
type Agregator struct {
	Count int
	Datas []string
}

var footprints map[[20]byte]*Agregator

// Agregators contains a map of Agregator
type Agregators []*Agregator

// NewAgregators create agregators
func NewAgregators() *Agregators {
	footprints = make(map[[20]byte]*Agregator, 0)

	return &Agregators{}
}

// Agregate agregate tokens according to positions
func (a *Agregators) Agregate(positions []int, tokens *[]sequence.Token, trans *transformer.Transformers) {
	var accumulator string
	var datas []string

	for _, i := range positions {
		result := trans.Apply(i, (*tokens)[i].Value)
		datas = append(datas, result)
		accumulator = accumulator + result
	}

	footprint := sha1.Sum([]byte(accumulator))

	if _, ok := footprints[footprint]; ok {
		footprints[footprint].Count = footprints[footprint].Count + 1
	} else {
		agregator := &Agregator{
			Count: 1,
			Datas: datas,
		}

		footprints[footprint] = agregator
		*a = append(*a, agregator)
	}
}

// ExtractPositions split positions fields from string
func ExtractPositions(fields string, size int) ([]int, error) {
	var positions []int

	if fields != "" {
		positionDups := make(map[int]bool, 0)

		for _, value := range strings.Split(fields, ",") {
			if position, err := strconv.Atoi(value); err == nil {
				if _, ok := positionDups[position]; ok == true {
					return []int{}, fmt.Errorf("This element is duplicated : %d", position)
				}

				if position >= size {
					return []int{}, fmt.Errorf("Position %d is greater or equal than maximum position %d", position, size)
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
