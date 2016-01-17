package agregator

import (
	"github.com/antham/goller/tokenizer"
	"github.com/antham/goller/transformer"
	"reflect"
	"testing"
)

func TestAgregateSingleValue(t *testing.T) {
	agregators := NewAgregators()

	tokens := []tokenizer.Token{
		{
			Value: "test1",
		},
		{
			Value: "test2",
		},
		{
			Value: "test3",
		},
		{
			Value: "test4",
		},
		{
			Value: "test5",
		},
		{
			Value: "test6",
		},
	}

	agregators.Agregate([]int{0, 1, 3}, &tokens, nil)

	for _, agregator := range *agregators {
		if agregator.Count != 1 {
			t.Error("Count must be 1 got", agregator.Count)
		}

		if !reflect.DeepEqual(agregator.Datas, []string{"test1", "test2", "test4"}) {
			t.Errorf("Count must be %s got %s", []string{"test1", "test2", "test4"}, agregator.Datas)
		}
	}
}

func TestAgregateSeveralValues(t *testing.T) {
	agregators := NewAgregators()

	for i := 0; i < 10; i++ {
		tokens := []tokenizer.Token{
			{
				Value: "test1",
			},
			{
				Value: "test2",
			},
			{
				Value: "test3",
			},
			{
				Value: "test4",
			},
			{
				Value: "test5",
			},
			{
				Value: "test6",
			},
		}

		if i <= 2 {
			agregators.Agregate([]int{0, 1}, &tokens, nil)
		} else if i > 2 && i <= 5 {
			agregators.Agregate([]int{2, 3}, &tokens, nil)
		} else if i > 5 {
			agregators.Agregate([]int{4, 5}, &tokens, nil)
		}
	}

	datas := *agregators

	if len(datas) != 3 {
		t.Error("Length must be 3 got", len(datas))
	}

	if datas[0].Count != 3 {
		t.Errorf("Count for key 0 must be 3 got %d", datas[0].Count)
	}

	if datas[1].Count != 3 {
		t.Errorf("Count for key 1 must be 3 got %d", datas[1].Count)
	}

	if datas[2].Count != 4 {
		t.Errorf("Count for key 2 must be 4 got %d", datas[2].Count)
	}
}

func TestApplyPreTransformer(t *testing.T) {
	agregators := NewAgregators()

	for i := 0; i < 10; i++ {
		tokens := []tokenizer.Token{
			{
				Value: "test1",
			},
			{
				Value: "test2",
			},
			{
				Value: "test3",
			},
		}

		trans := transformer.NewTransformers()
		trans.Append(0, "upp", []string{})
		trans.Append(2, "upp", []string{})

		agregators.Agregate(
			[]int{0, 1, 2},
			&tokens,
			trans,
		)
	}

	datas := *agregators

	if len(datas) != 1 {
		t.Error("Length must be 1 got", len(datas))
	}

	if datas[0].Count != 10 {
		t.Errorf("Count for key 0 must be 10, got %d", datas[0].Count)
	}
}
