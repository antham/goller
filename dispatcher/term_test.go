package dispatcher

import (
	"bytes"
	"github.com/antham/goller/agregator"
	"github.com/antham/goller/transformer"
	"github.com/trustpath/sequence"
	"io"
	"os"
	"testing"
)

func TestTermDispatcherDisplaying(t *testing.T) {
	agregators := agregator.NewAgregators()

	tokens := []sequence.Token{
		sequence.Token{
			Value: "test1",
		},
		sequence.Token{
			Value: "test2",
		},
		sequence.Token{
			Value: "test3",
		},
		sequence.Token{
			Value: "test4",
		},
		sequence.Token{
			Value: "test5",
		},
		sequence.Token{
			Value: "test6",
		},
	}

	agregators.Agregate([]int{0, 1, 3}, &tokens, transformer.TransformersMap{})

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	d := NewTermDispatcher("|")
	d.RenderItems(agregators.Get())

	w.Close()
	os.Stdout = old
	out := <-outC

	if out != "1|test1|test2|test4\n" {
		t.Errorf("Must output %s got %s", "1|test1|test2|test4\n", out)
	}
}
