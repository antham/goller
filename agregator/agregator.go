package agregator

import (
	"crypto/sha1"
	"github.com/antham/goller/tokenizer"
	"github.com/antham/goller/transformer"
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
func (a *Agregators) Agregate(positions []int, tokens *[]tokenizer.Token, trans *transformer.Transformers) {
	var accumulator string
	var datas []string

	for _, i := range positions {
		result := (*tokens)[i].Value

		if trans != nil {
			result = trans.Apply(i, result)
		}

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
