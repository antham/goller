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
