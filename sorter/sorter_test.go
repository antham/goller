package sorter

import (
	"github.com/antham/goller/agregator"
	"testing"
)

func TestInt(t *testing.T) {
	datas1 := []string{"5", "8"}
	datas2 := []string{"4", "9"}
	datas3 := []string{"3", "1"}

	agregators := agregator.Agregators{
		0: &agregator.Agregator{
			Count: 5,
			Datas: datas1,
			DatasOrdered: map[int]*string{
				0: &datas1[0],
				1: &datas1[1],
			},
		},
		1: &agregator.Agregator{
			Count: 5,
			Datas: datas2,
			DatasOrdered: map[int]*string{
				0: &datas2[0],
				1: &datas2[1],
			},
		},
		2: &agregator.Agregator{
			Count: 5,
			Datas: datas3,
			DatasOrdered: map[int]*string{
				0: &datas3[0],
				1: &datas3[1],
			},
		},
	}

	sorters := NewSorters()
	sorters.Append(-1, 1, "int", []string{})
	sorters.Sort(&agregators)

	if *agregators[0].DatasOrdered[1] != "1" ||
		*agregators[1].DatasOrdered[1] != "8" ||
		*agregators[2].DatasOrdered[1] != "9" {
		t.Errorf("Got %s,%s,%s, expected order 1,8,9", *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[2])
	}
}

func TestStrl(t *testing.T) {
	datas1 := []string{"5", "hello"}
	datas2 := []string{"4", "everybody"}
	datas3 := []string{"3", "!"}

	agregators := agregator.Agregators{
		0: &agregator.Agregator{
			Count: 5,
			Datas: datas1,
			DatasOrdered: map[int]*string{
				0: &datas1[0],
				1: &datas1[1],
			},
		},
		1: &agregator.Agregator{
			Count: 5,
			Datas: datas2,
			DatasOrdered: map[int]*string{
				0: &datas2[0],
				1: &datas2[1],
			},
		},
		2: &agregator.Agregator{
			Count: 5,
			Datas: datas3,
			DatasOrdered: map[int]*string{
				0: &datas3[0],
				1: &datas3[1],
			},
		},
	}

	sorters := NewSorters()
	sorters.Append(-1, 1, "strl", []string{})
	sorters.Sort(&agregators)

	if *agregators[0].DatasOrdered[1] != "!" ||
		*agregators[1].DatasOrdered[1] != "hello" ||
		*agregators[2].DatasOrdered[1] != "everybody" {
		t.Errorf("Got %s,%s,%s, expected order !,hello,everybody", *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[2])
	}
}

func TestStr(t *testing.T) {
	datas1 := []string{"5", "hello"}
	datas2 := []string{"4", "everybody"}
	datas3 := []string{"3", "!"}

	agregators := agregator.Agregators{
		0: &agregator.Agregator{
			Count: 5,
			Datas: datas1,
			DatasOrdered: map[int]*string{
				0: &datas1[0],
				1: &datas1[1],
			},
		},
		1: &agregator.Agregator{
			Count: 5,
			Datas: datas2,
			DatasOrdered: map[int]*string{
				0: &datas2[0],
				1: &datas2[1],
			},
		},
		2: &agregator.Agregator{
			Count: 5,
			Datas: datas3,
			DatasOrdered: map[int]*string{
				0: &datas3[0],
				1: &datas3[1],
			},
		},
	}

	sorters := NewSorters()
	sorters.Append(-1, 1, "str", []string{})
	sorters.Sort(&agregators)

	if *agregators[0].DatasOrdered[1] != "!" ||
		*agregators[1].DatasOrdered[1] != "everybody" ||
		*agregators[2].DatasOrdered[1] != "hello" {
		t.Errorf("Got %s,%s,%s, expected order !,everybody,hello", *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[2])
	}
}
