package cli

import (
	"github.com/antham/goller/agregator"
	"github.com/antham/goller/dispatcher"
	"github.com/antham/goller/reader"
	"github.com/antham/goller/tokenizer"
	"gopkg.in/alecthomas/kingpin.v2"
	"reflect"
	"strings"
	"testing"
)

func TestGroup(t *testing.T) {

	input := strings.NewReader("hello world !\nhello world !\nHi everybody !")
	r := reader.Reader{
		Input: input,
	}

	switch kingpin.MustParse(app.Parse(strings.Fields("group whi 0,1,2,3"))) {
	case cmd["group"].FullCommand():
		group := &group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			args:       groupArgs,
		}

		group.Consume()

		agregators := group.agrBuilder.Get()

		if len(*agregators) != 2 {
			t.Errorf("Got %d length, expected %d", len(*agregators), 2)
		}

		expected := []string{
			"2",
			"hello",
			"world",
			"!",
		}

		got := (*agregators)[0].Datas

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		expected = []string{
			"1",
			"Hi",
			"everybody",
			"!",
		}

		got = (*agregators)[1].Datas

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}
}
