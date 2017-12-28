package configurator

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

// Handler handle configurators
type Handler struct {
	filename string
}

// NewHandler create a new handler
func NewHandler(filename string) (*Handler, error) {
	baseDir, err := user.Current()

	if err != nil {
		return nil, err
	}

	return &Handler{
		filename: path.Join(baseDir.HomeDir, filename),
	}, nil
}

// Load create a configurator from saved config file
func (h *Handler) Load() (*Configurator, error) {
	content, err := ioutil.ReadFile(h.filename)

	c := New()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Save configurator to config file
func (h *Handler) Save(config *Configurator) error {
	configJSON, err := json.Marshal(config)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(h.filename, configJSON, os.ModePerm)
}
