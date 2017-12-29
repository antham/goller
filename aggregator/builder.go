package aggregator

import (
	"github.com/antham/goller/tokenizer"
	"github.com/antham/goller/transformer"
	"strconv"
)

// Builder wraps operations needed to manage aggregators
type Builder struct {
	aggregators *Aggregators
	footprints  map[string]*Aggregator
}

// NewBuilder create builder
func NewBuilder() *Builder {
	return &Builder{
		NewAggregators(),
		map[string]*Aggregator{},
	}
}

// iterate create a new aggregator or increase existing aggregator counter
func (b *Builder) iterate(datas []*string, datasOrdered map[int]*string, accumulator string) {
	if _, ok := (*b).footprints[accumulator]; ok {
		(*b).footprints[accumulator].Count = (*b).footprints[accumulator].Count + 1
	} else {
		aggregator := &Aggregator{
			Count:        1,
			DatasOrdered: datasOrdered,
			Datas:        datas,
		}

		(*b).footprints[accumulator] = aggregator
		(*b.aggregators) = append((*b.aggregators), aggregator)
	}
}

// Aggregate aggregates tokens according to positions
func (b *Builder) Aggregate(positions []int, tokens *[]tokenizer.Token, trans *transformer.Transformers) {
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
	for _, aggregator := range *b.aggregators {
		if _, ok := (*aggregator).DatasOrdered[0]; ok {
			*((*aggregator).DatasOrdered[0]) = strconv.Itoa((*aggregator).Count)
		}
	}
}

// Get retrieve aggregators from builder
func (b *Builder) Get() *Aggregators {
	return b.aggregators
}
