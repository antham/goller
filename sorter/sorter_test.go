package sorter

import (
	"github.com/antham/goller/agregator"
	"github.com/antham/goller/tokenizer"
	"testing"
)

func TestIntMultiSort(t *testing.T) {
	datas := [][]string{
		{"3", "8", "2"},
		{"4", "9", "3"},
		{"3", "8", "0"},
		{"3", "1", "10"},
		{"3", "9", "1"},
		{"1", "9", "1"},
		{"2", "9", "1"},
	}

	builder := agregator.NewBuilder()

	for _, data := range datas {
		tokens := []tokenizer.Token{}

		for _, item := range data {
			tokens = append(tokens, tokenizer.Token{Value: item})
		}

		builder.Aggregate([]int{1, 2, 3}, &tokens, nil)
	}

	builder.SetCounterIfAny()
	agregators := builder.Get()

	sorters := NewSorters()
	sorters.Append(1, "int", []string{})
	sorters.Append(2, "int", []string{})
	sorters.Append(3, "int", []string{})
	sorters.Sort(agregators)

	expected := [][]*string{}

	for i, row := range [][]string{
		{"1", "9", "1"},
		{"2", "9", "1"},
		{"3", "1", "10"},
		{"3", "8", "0"},
		{"3", "8", "2"},
		{"3", "9", "1"},
		{"4", "9", "3"},
	} {
		expected = append(expected, []*string{})

		for _, element := range row {
			e := element
			expected[i] = append(expected[i], &e)
		}
	}

	for i := 0; i < len(expected); i++ {
		for j := 0; j < 3; j++ {
			if *(*agregators)[i].Datas[j] != *expected[i][j] {
				t.Errorf("Got %s, expected %s", *(*agregators)[i].Datas[j], *expected[i][j])
			}
		}
	}
}

func TestInt(t *testing.T) {
	var datas1 []*string
	var datas2 []*string
	var datas3 []*string

	for _, row := range [][]string{
		{"5", "4", "3"},
		{"8", "9", "1"},
	} {
		datas1 = append(datas1, &row[0])
		datas2 = append(datas2, &row[1])
		datas3 = append(datas3, &row[2])
	}

	agregators := agregator.Agregators{
		0: &agregator.Agregator{
			Count: 5,
			Datas: datas1,
			DatasOrdered: map[int]*string{
				0: datas1[0],
				1: datas1[1],
			},
		},
		1: &agregator.Agregator{
			Count: 5,
			Datas: datas2,
			DatasOrdered: map[int]*string{
				0: datas2[0],
				1: datas2[1],
			},
		},
		2: &agregator.Agregator{
			Count: 5,
			Datas: datas3,
			DatasOrdered: map[int]*string{
				0: datas3[0],
				1: datas3[1],
			},
		},
	}

	sorters := NewSorters()
	sorters.Append(1, "int", []string{})
	sorters.Sort(&agregators)

	if *agregators[0].DatasOrdered[1] != "1" ||
		*agregators[1].DatasOrdered[1] != "8" ||
		*agregators[2].DatasOrdered[1] != "9" {
		t.Errorf("Got %s,%s,%s, expected order 1,8,9", *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[2])
	}
}

func TestStrl(t *testing.T) {
	var datas1 []*string
	var datas2 []*string
	var datas3 []*string

	for _, row := range [][]string{
		{"5", "4", "3"},
		{"hello", "everybody", "!"},
	} {
		datas1 = append(datas1, &row[0])
		datas2 = append(datas2, &row[1])
		datas3 = append(datas3, &row[2])
	}

	agregators := agregator.Agregators{
		0: &agregator.Agregator{
			Count: 5,
			Datas: datas1,
			DatasOrdered: map[int]*string{
				0: datas1[0],
				1: datas1[1],
			},
		},
		1: &agregator.Agregator{
			Count: 5,
			Datas: datas2,
			DatasOrdered: map[int]*string{
				0: datas2[0],
				1: datas2[1],
			},
		},
		2: &agregator.Agregator{
			Count: 5,
			Datas: datas3,
			DatasOrdered: map[int]*string{
				0: datas3[0],
				1: datas3[1],
			},
		},
	}

	sorters := NewSorters()
	sorters.Append(1, "strl", []string{})
	sorters.Sort(&agregators)

	if *agregators[0].DatasOrdered[1] != "!" ||
		*agregators[1].DatasOrdered[1] != "hello" ||
		*agregators[2].DatasOrdered[1] != "everybody" {
		t.Errorf("Got %s,%s,%s, expected order !,hello,everybody", *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[2])
	}
}

func TestStr(t *testing.T) {
	var datas1 []*string
	var datas2 []*string
	var datas3 []*string

	for _, row := range [][]string{
		{"5", "4", "3"},
		{"hello", "everybody", "!"},
	} {
		datas1 = append(datas1, &row[0])
		datas2 = append(datas2, &row[1])
		datas3 = append(datas3, &row[2])
	}

	agregators := agregator.Agregators{
		0: &agregator.Agregator{
			Count: 5,
			Datas: datas1,
			DatasOrdered: map[int]*string{
				0: datas1[0],
				1: datas1[1],
			},
		},
		1: &agregator.Agregator{
			Count: 5,
			Datas: datas2,
			DatasOrdered: map[int]*string{
				0: datas2[0],
				1: datas2[1],
			},
		},
		2: &agregator.Agregator{
			Count: 5,
			Datas: datas3,
			DatasOrdered: map[int]*string{
				0: datas3[0],
				1: datas3[1],
			},
		},
	}

	sorters := NewSorters()
	sorters.Append(1, "str", []string{})
	sorters.Sort(&agregators)

	if *agregators[0].DatasOrdered[1] != "!" ||
		*agregators[1].DatasOrdered[1] != "everybody" ||
		*agregators[2].DatasOrdered[1] != "hello" {
		t.Errorf("Got %s,%s,%s, expected order !,everybody,hello", *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[1], *agregators[0].DatasOrdered[2])
	}
}
