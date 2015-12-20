package agregator

import (
	"crypto/sha1"
	"fmt"
	"github.com/trustpath/sequence"
	"strconv"
	"strings"
)

//Agregator represents a unique log line
type Agregator struct {
	Count int
	Datas []string
}

//Agregators contains a map of Agregator
type Agregators struct {
	agregators map[[20]byte]*Agregator
}

//NewAgregators create agregators
func NewAgregators() *Agregators {
	return &Agregators{
		agregators: make(map[[20]byte]*Agregator, 0),
	}
}

//Agregate agregate tokens acoording to positions
func (a *Agregators) Agregate(positions []int, tokens *[]sequence.Token) {
	var accumulator string
	var datas []string

	for _, i := range positions {
		datas = append(datas, (*tokens)[i].Value)
		accumulator = accumulator + (*tokens)[i].Value
	}

	footprint := sha1.Sum([]byte(accumulator))

	if _, ok := a.agregators[footprint]; ok {
		a.agregators[footprint].Count = a.agregators[footprint].Count + 1
	} else {
		a.agregators[footprint] = &Agregator{
			Count: 1,
			Datas: datas,
		}
	}
}

//Get return agregated values
func (a *Agregators) Get() map[[20]byte]*Agregator {
	return a.agregators
}

//ExtractPositions split positions fields from string
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
