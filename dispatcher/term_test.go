package dispatcher

import (
	"bytes"
	"github.com/antham/goller/v2/aggregator"
	"github.com/antham/goller/v2/tokenizer"
	"github.com/antham/goller/v2/transformer"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTermDispatcherDisplaying(t *testing.T) {
	builder := aggregator.NewBuilder()

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

	builder.Aggregate([]int{0, 1, 3}, &tokens, transformer.NewTransformers())
	builder.SetCounterIfAny()

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		assert.NoError(t, err)
		outC <- buf.String()
	}()

	d := NewTermDispatch("|")
	d.RenderItems(builder.Get())

	w.Close()
	os.Stdout = old
	out := <-outC

	if out != "1|test1|test3\n" {
		t.Errorf("Must output %s got %s", "1|test1|test2|test4\n", out)
	}
}
