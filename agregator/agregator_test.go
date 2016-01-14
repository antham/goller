package agregator

import (
	"crypto/sha1"
	"fmt"
	"github.com/antham/goller/transformer"
	"github.com/trustpath/sequence"
	"reflect"
	"testing"
)

func TestExtractPositionsFromString(t *testing.T) {
	positions, err := ExtractPositions("2,5,0,8,16,12", 17)

	if err != nil {
		t.Error("Got an error", err)
	}

	if len(positions) != 6 {
		t.Error("Expected slice length of 6, got", len(positions))
	}

	if positions[0] != 2 {
		t.Error("First element must be 2, got", positions[0])
	}

	if positions[5] != 12 {
		t.Error("First element must be 12, got", positions[5])
	}
}

func TestExtractPositionsMustReturnUniquePositions(t *testing.T) {
	_, err := ExtractPositions("2,5,5,8,16,12", 17)

	if err == nil || err.Error() != "This element is duplicated : 5" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromEmptyString(t *testing.T) {
	_, err := ExtractPositions("", 17)

	if err == nil || err.Error() != "At least 1 element is required" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromStringContainingSomethingDifferentThanNumber(t *testing.T) {
	_, err := ExtractPositions("1,2,a,b,c", 17)

	if err == nil || err.Error() != "a is not a number" {
		t.Error("An error must occur", err)
	}
}

func TestExtractPositionsFromStringContainingPositionOverLimit(t *testing.T) {
	_, err := ExtractPositions("1,2,5,11,8,9", 11)

	if err == nil || err.Error() != "Position 11 is greater or equal than maximum position 11" {
		t.Error("An error must occur", err)
	}

	_, err = ExtractPositions("1,2,5,12,8,9", 11)

	if err == nil || err.Error() != "Position 12 is greater or equal than maximum position 11" {
		t.Error("An error must occur", err)
	}
}

func TestAgregateSingleValue(t *testing.T) {
	agregators := NewAgregators()

	tokens := []sequence.Token{
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

	agregators.Agregate([]int{0, 1, 3}, &tokens, transformer.NewTransformers())

	for key, agregator := range agregators.Get() {
		id := fmt.Sprintf("%x", key)

		if id != "3f5b45cb95d336f2c9449ee6e716f736f24e6fd1" {
			t.Error("Id must be 3f5b45cb95d336f2c9449ee6e716f736f24e6fd1 got", id)
		}

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
		tokens := []sequence.Token{
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
			agregators.Agregate([]int{0, 1}, &tokens, transformer.NewTransformers())
		} else if i > 2 && i <= 5 {
			agregators.Agregate([]int{2, 3}, &tokens, transformer.NewTransformers())
		} else if i > 5 {
			agregators.Agregate([]int{4, 5}, &tokens, transformer.NewTransformers())
		}
	}

	datas := agregators.Get()

	if len(datas) != 3 {
		t.Error("Length must be 3 got", len(datas))
	}

	key := sha1.Sum([]byte("test1" + "test2"))

	if datas[key].Count != 3 {
		t.Errorf("Count for key %x must be 3 got %d", key, datas[key].Count)
	}

	key = sha1.Sum([]byte("test3" + "test4"))

	if datas[key].Count != 3 {
		t.Errorf("Count for key %x must be 3 got %d", key, datas[key].Count)
	}

	key = sha1.Sum([]byte("test5" + "test6"))

	if datas[key].Count != 4 {
		t.Errorf("Count for key %x must be 4 got %d", key, datas[key].Count)
	}
}

func TestApplyPreTransformer(t *testing.T) {
	agregators := NewAgregators()

	for i := 0; i < 10; i++ {
		tokens := []sequence.Token{
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

	datas := agregators.Get()

	if len(datas) != 1 {
		t.Error("Length must be 1 got", len(datas))
	}

	key := sha1.Sum([]byte("TEST1" + "test2" + "TEST3"))

	if _, ok := datas[key]; ok != true {
		t.Error("Entry for TEST1 test2 TEST3 doesn't exist")
	}

	if datas[key].Count != 10 {
		t.Errorf("Count for key %x must be 10, got %d", key, datas[key].Count)
	}
}
