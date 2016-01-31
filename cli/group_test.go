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
	app := initApp()
	cmd := initCmd(app)
	groupArgs := initGroupArgs(cmd["group"])

	positions := []int{0, 1, 2, 3}

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
			positions:  &positions,
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

func TestGroupWithTransformers(t *testing.T) {
	app := initApp()
	cmd := initCmd(app)
	groupArgs := initGroupArgs(cmd["group"])

	positions := []int{0, 1, 2, 3}

	input := strings.NewReader("1 2 3\n1 2 3\n1 2 3")
	r := reader.Reader{
		Input: input,
	}

	switch kingpin.MustParse(app.Parse([]string{"group", "whi", "-t", `1:add("1")`, "-t", `2:add("2")`, "-t", `3:add("3")`, `0,1,2,3`})) {
	case cmd["group"].FullCommand():
		group := &group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			positions:  &positions,
			args:       groupArgs,
		}

		group.Consume()

		agregators := group.agrBuilder.Get()

		if len(*agregators) != 1 {
			t.Errorf("Got %d length, expected %d", len(*agregators), 1)
		}

		expected := []string{
			"3",
			"2",
			"4",
			"6",
		}

		got := (*agregators)[0].Datas

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}
}

func TestGroupWithSorters(t *testing.T) {
	app := initApp()
	cmd := initCmd(app)
	groupArgs := initGroupArgs(cmd["group"])

	positions := []int{0, 1, 2, 3}

	input := strings.NewReader("3 2 1\n6 5 4\n9 8 7")
	r := reader.Reader{
		Input: input,
	}

	switch kingpin.MustParse(app.Parse([]string{"group", "whi", "-s", `3:int,2:int,1:int`, `0,1,2,3`})) {
	case cmd["group"].FullCommand():
		group := &group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			positions:  &positions,
			args:       groupArgs,
		}

		group.Consume()

		agregators := group.agrBuilder.Get()

		if len(*agregators) != 3 {
			t.Errorf("Got %d length, expected %d", len(*agregators), 1)
		}

		expected := []string{
			"1",
			"3",
			"2",
			"1",
		}

		got := (*agregators)[0].Datas

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		expected = []string{
			"1",
			"9",
			"8",
			"7",
		}

		got = (*agregators)[2].Datas

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}
}
