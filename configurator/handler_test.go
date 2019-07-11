package configurator

import (
	"io/ioutil"
	"os/user"
	"path"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createConfig(filename string, data string, t *testing.T) {
	baseDir, err := user.Current()

	if err != nil {
		t.Error("Can't retrieve user home path")
	}

	buf := []byte(data)
	assert.NoError(t, ioutil.WriteFile(path.Join(baseDir.HomeDir, filename), buf, 0666))
}

func readConfig(filename string, t *testing.T) ([]byte, error) {
	baseDir, err := user.Current()

	if err != nil {
		t.Error("Can't retrieve user home path")
	}

	return ioutil.ReadFile(path.Join(baseDir.HomeDir, filename))
}

func TestLoadUnexistentFile(t *testing.T) {
	handler, err := NewHandler("whatever")

	if err != nil {
		t.Error("Can't create an handler")
	}

	_, err = handler.Load()

	if reflect.TypeOf(err).String() != "*os.PathError" {
		t.Error("Must return a path error")
	}
}

func TestLoadInvalidJSON(t *testing.T) {
	filename := "test.json"

	createConfig(filename, "test", t)

	handler, _ := NewHandler(filename)

	_, err := handler.Load()

	if reflect.TypeOf(err).String() != "*json.SyntaxError" {
		t.Error("Must return a json syntax error")
	}
}

func TestLoadValidJSON(t *testing.T) {
	filename := "test.json"

	createConfig(filename, "{}", t)

	handler, _ := NewHandler(filename)

	c, err := handler.Load()

	if err != nil {
		t.Error("Must return no error")
	}

	if reflect.DeepEqual(c.Bindings.Sorters, &map[string]string{}) {
		t.Error("Sorters must be an empty map")
	}

	if reflect.DeepEqual(c.Bindings.Transformers, &map[string]string{}) {
		t.Error("Transformers must be an empty map")
	}
}

func TestSaveConfig(t *testing.T) {
	filename := "test.json"

	handler, _ := NewHandler(filename)

	category := "bindings.transformers"
	key := "hello"
	value := "world"

	c := New()
	assert.NoError(t, c.Set(&category, &key, &value))

	err := handler.Save(c)

	if err != nil {
		t.Error("Must return no error")
	}

	jsonConfig, err := readConfig(filename, t)

	if err != nil {
		t.Error("Must return no error")
	}

	if string(jsonConfig) != `{"Bindings":{"Transformers":{"hello":"world"},"Sorters":{}}}` {
		t.Error("Must return no error")
	}
}

func TestSaveConfigWithUnexistingFilename(t *testing.T) {
	filename := "test.json"

	handler, _ := NewHandler(filename)
	handler.filename = "/whatever/whatever"

	c := New()

	err := handler.Save(c)

	if err.Error() != "open /whatever/whatever: no such file or directory" {
		t.Error("Must return an error")
	}
}
