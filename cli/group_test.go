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
		group := &Group{
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

		test1 := "2"
		test2 := "hello"
		test3 := "world"
		test4 := "!"

		expected := []*string{
			&test1,
			&test2,
			&test3,
			&test4,
		}

		got := (*agregators)[0].Datas

		for i := 0; i < len(got); i++ {
			if !reflect.DeepEqual(got[i], expected[i]) {
				t.Errorf("Got %s, expected %s", *got[i], *expected[i])
			}
		}

		test1 = "1"
		test2 = "Hi"
		test3 = "everybody"
		test4 = "!"

		expected = []*string{
			&test1,
			&test2,
			&test3,
			&test4,
		}

		got = (*agregators)[1].Datas

		for i := 0; i < len(got); i++ {
			if !reflect.DeepEqual(got[i], expected[i]) {
				t.Errorf("Got %s, expected %s", *got[i], *expected[i])
			}
		}
	}
}

func TestGroupWithDifferenteSizeLines(t *testing.T) {
	app := initApp()
	cmd := initCmd(app)
	groupArgs := initGroupArgs(cmd["group"])

	triggerFatalError = func(err error) {
		if err == nil || err.Error() != "Wrong parsing strategy (based on first line tokenization), got 1 tokens instead of 3\nLine : 1\n" {
			t.Error("Must return an error, lines are not parsed with the same number of tokens")
		}
	}

	positions := []int{1}

	input := strings.NewReader("1 2 3\n1\n1 2 3 4 5\n4 5 6")
	r := reader.Reader{
		Input: input,
	}

	ignore := false

	switch kingpin.MustParse(app.Parse([]string{"group", "whi", "1"})) {
	case cmd["group"].FullCommand():
		group := &Group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			positions:  &positions,
			ignore:     &ignore,
			args:       groupArgs,
		}

		group.Consume()
	}
}

func TestGroupWithNoTokenParsed(t *testing.T) {
	app := initApp()
	cmd := initCmd(app)
	groupArgs := initGroupArgs(cmd["group"])

	triggerFatalError = func(err error) {
		if err == nil || err.Error() != "No tokens found" {
			t.Error("Must return an error, no tokens are extracted")
		}
	}

	positions := []int{1}

	input := strings.NewReader("1 2 3")
	r := reader.Reader{
		Input: input,
	}

	ignore := false

	switch kingpin.MustParse(app.Parse([]string{"group", `reg("whatever")`, "1"})) {
	case cmd["group"].FullCommand():
		group := &Group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			positions:  &positions,
			ignore:     &ignore,
			args:       groupArgs,
		}

		group.Consume()
	}
}

func TestGroupWithAPositionGreaterThanExistingPosition(t *testing.T) {
	app := initApp()
	cmd := initCmd(app)
	groupArgs := initGroupArgs(cmd["group"])

	triggerFatalError = func(err error) {
		if err == nil || err.Error() != "A position is greater or equal than maximum position 3" {
			t.Error("Must return an error, position is greater than max position")
		}
	}

	positions := []int{1, 2, 3, 4}

	input := strings.NewReader("1 2 3")
	r := reader.Reader{
		Input: input,
	}

	ignore := false

	switch kingpin.MustParse(app.Parse([]string{"group", `whi`, "1,2,3,4"})) {
	case cmd["group"].FullCommand():
		group := &Group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			positions:  &positions,
			ignore:     &ignore,
			args:       groupArgs,
		}

		group.Consume()
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
		group := &Group{
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

		test1 := "3"
		test2 := "2"
		test3 := "4"
		test4 := "6"

		expected := []*string{
			&test1,
			&test2,
			&test3,
			&test4,
		}

		got := (*agregators)[0].Datas

		for i := 0; i < len(got); i++ {
			if !reflect.DeepEqual(got[i], expected[i]) {
				t.Errorf("Got %s, expected %s", *got[i], *expected[i])
			}
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
		group := &Group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			positions:  &positions,
			args:       groupArgs,
		}

		group.Consume()
		group.Sort()

		agregators := group.agrBuilder.Get()

		if len(*agregators) != 3 {
			t.Errorf("Got %d length, expected %d", len(*agregators), 1)
		}

		test1 := "1"
		test2 := "3"
		test3 := "2"
		test4 := "1"

		expected := []*string{
			&test1,
			&test2,
			&test3,
			&test4,
		}

		got := (*agregators)[0].Datas

		for i := 0; i < len(got); i++ {
			if !reflect.DeepEqual(got[i], expected[i]) {
				t.Errorf("Got %s, expected %s", *got[i], *expected[i])
			}
		}

		test1 = "1"
		test2 = "9"
		test3 = "8"
		test4 = "7"

		expected = []*string{
			&test1,
			&test2,
			&test3,
			&test4,
		}

		got = (*agregators)[2].Datas

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}
}

func TestGroupWithIgnoreFlag(t *testing.T) {
	app := initApp()
	cmd := initCmd(app)
	groupArgs := initGroupArgs(cmd["group"])

	positions := []int{1}

	input := strings.NewReader("1 2 3\n1\n1 2 3 4 5\n4 5 6")
	r := reader.Reader{
		Input: input,
	}

	ignore := true

	switch kingpin.MustParse(app.Parse([]string{"group", "-i", "whi", "1"})) {
	case cmd["group"].FullCommand():
		group := &Group{
			tokenizer:  *tokenizer.NewTokenizer(groupArgs.parser.Get()),
			agrBuilder: *agregator.NewBuilder(),
			dispatcher: dispatcher.NewTermDispatch(*groupArgs.delimiter),
			reader:     r,
			positions:  &positions,
			ignore:     &ignore,
			args:       groupArgs,
		}

		group.Consume()

		agregators := group.agrBuilder.Get()

		if len(*agregators) != 2 {
			t.Errorf("Got %d length, expected %d", len(*agregators), 1)
		}

		test1 := "1"
		test2 := "4"

		expected := []*string{
			&test1,
			&test2,
		}

		got := (*agregators)[0].Datas

		for i := 0; i < len(got); i++ {
			if !reflect.DeepEqual(got[i], expected[i]) {
				t.Errorf("Got %s, expected %s", *got[i], *expected[i])
			}
		}
	}
}
