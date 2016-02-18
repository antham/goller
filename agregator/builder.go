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

// iterate create a new agregator or increase existing agregator counter
func (b *Builder) iterate(datas []*string, datasOrdered map[int]*string, accumulator string) {
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
}

// Agregate agregate tokens according to positions
func (b *Builder) Agregate(positions []int, tokens *[]tokenizer.Token, trans *transformer.Transformers) {
	var accumulator string
	datas := []*string{}
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
		}

		datas = append(datas, &result)
		datasOrdered[i] = &result
	}

	b.iterate(datas, datasOrdered, accumulator)
}

// SetCounterIfAny set counter value among other value fields
func (b *Builder) SetCounterIfAny() {
	for _, agregator := range *b.agregators {
		if _, ok := (*agregator).DatasOrdered[0]; ok == true {
			*((*agregator).DatasOrdered[0]) = strconv.Itoa((*agregator).Count)
		}
	}
}

// Get retrieve agregators from builder
func (b *Builder) Get() *Agregators {
	return b.agregators
}
