package cli

import (
	"log"

	"github.com/antham/goller/v2/reader"
	"github.com/antham/goller/v2/tokenizer"
	"gopkg.in/alecthomas/kingpin.v2"
	"reflect"
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {
	app := initApp()
	cmd := initCmd(app)
	tokenizeArgs := initTokenizeArgs(cmd["tokenize"])

	input := strings.NewReader("hello world !\nhello world !\nHi everybody !")
	r := reader.Reader{
		Input: input,
	}

	switch kingpin.MustParse(app.Parse(strings.Fields("tokenize whi"))) {
	case cmd["group"].FullCommand():

		tokenize := &Tokenize{
			tokenizer: *tokenizer.NewTokenizer(tokenizeArgs.parser.Get()),
			reader:    r,
		}

		err := tokenize.Tokenize()

		if err != nil {
			log.Fatal(err)
		}

		if len(tokenize.tokens) != 3 {
			t.Errorf("Got %d length, expected %d", len(tokenize.tokens), 2)
		}

		got := []tokenizer.Token{
			{
				Value: "hello",
			},
			{
				Value: "world",
			},
			{
				Value: "!",
			},
		}

		for i := 0; i < len(got); i++ {
			if !reflect.DeepEqual(got[i].Value, tokenize.tokens[i].Value) {
				t.Errorf("Got %s, expected %s", got[i].Value, tokenize.tokens[i].Value)
			}
		}
	}
}
