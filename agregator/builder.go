package agregator

import (
	"crypto/sha1"
	"github.com/antham/goller/tokenizer"
	"github.com/antham/goller/transformer"
	"strconv"
)

// Builder wraps operations needed to manage agregators
type Builder struct {
	agregators *Agregators
	footprints map[[20]byte]*Agregator
}

// NewBuilder create builder
func NewBuilder() *Builder {
	return &Builder{
		NewAgregators(),
		make(map[[20]byte]*Agregator, 0),
	}
}

// Agregate agregate tokens according to positions
func (b *Builder) Agregate(positions []int, tokens *[]tokenizer.Token, trans *transformer.Transformers) {
	var accumulator string
	var counterIndex = -1
	datas := []string{}
	datasOrdered := map[int]*string{}

	for _, i := range positions {
		var result string

		if i > 0 {
			index := i - 1

			result = (*tokens)[index].Value

			if trans != nil {
				result = trans.Apply(i, result)
			}

			accumulator = accumulator + result
		} else {
			result = ""
			counterIndex = len(datas)
		}

		datas = append(datas, result)
		datasOrdered[i] = &datas[len(datas)-1]
	}

	footprint := sha1.Sum([]byte(accumulator))

	if _, ok := (*b).footprints[footprint]; ok {
		(*b).footprints[footprint].Count = (*b).footprints[footprint].Count + 1
	} else {
		agregator := &Agregator{
			Count:        1,
			DatasOrdered: datasOrdered,
			Datas:        datas,
		}

		(*b).footprints[footprint] = agregator
		(*b.agregators) = append((*b.agregators), agregator)
	}

	if counterIndex != -1 {
		(*b).footprints[footprint].Datas[counterIndex] = strconv.Itoa((*b).footprints[footprint].Count)
		(*b).footprints[footprint].DatasOrdered[counterIndex] = &(*b).footprints[footprint].Datas[counterIndex]
	}
}

// Get retrieve agregators from builder
func (b *Builder) Get() *Agregators {
	return b.agregators
}
