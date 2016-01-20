package agregator

import (
	"github.com/antham/goller/tokenizer"
	"github.com/antham/goller/transformer"
	"reflect"
	"testing"
)

func TestAgregateSingleTokenWithNoVisibleCounter(t *testing.T) {
	builder := NewBuilder()

	tokens := []tokenizer.Token{
		{
			Value: "test1",
		},
	}

	builder.Agregate([]int{1}, &tokens, nil)

	for _, agregator := range *builder.Get() {
		if agregator.Count != 1 {
			t.Error("Count must be 1 got", agregator.Count)
		}

		got := agregator.Datas
		expected := []string{"test1"}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %s, expected %s", got, expected)
		}
	}
}

func TestAgregateSingleToken(t *testing.T) {
	builder := NewBuilder()

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

	builder.Agregate([]int{3, 6, 0, 1}, &tokens, nil)

	for _, agregator := range *builder.Get() {
		if agregator.Count != 1 {
			t.Error("Count must be 1 got", agregator.Count)
		}

		got := agregator.Datas
		expected := []string{"test3", "test6", "1", "test1"}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %s, expected %s", got, expected)
		}
	}
}

func TestAgregateSeveralToken(t *testing.T) {
	builder := NewBuilder()

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
			builder.Agregate([]int{0, 1}, &tokens, nil)
		} else if i > 2 && i <= 5 {
			builder.Agregate([]int{2, 3}, &tokens, nil)
		} else if i > 5 {
			builder.Agregate([]int{4, 5}, &tokens, nil)
		}
	}

	datas := *builder.Get()

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
	builder := NewBuilder()

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
		trans.Append(1, "upp", []string{})
		trans.Append(2, "upp", []string{})

		builder.Agregate(
			[]int{0, 1, 2},
			&tokens,
			trans,
		)
	}

	datas := *builder.Get()

	if len(datas) != 1 {
		t.Error("Length must be 1 got", len(datas))
	}

	if datas[0].Count != 10 {
		t.Errorf("Count for key 0 must be 10, got %d", datas[0].Count)
	}

	got := datas[0].Datas
	expected := []string{"10", "TEST1", "TEST2"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got %v, expected %v", got, expected)
	}
}
