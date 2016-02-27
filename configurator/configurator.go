package configurator

import (
	"fmt"
	// "strings"
)

// Items represents a config file option
type Items map[string]string

// Configurator is a list of config entries
type Configurator struct {
	Bindings struct {
		Transformers *Items
		Sorters      *Items
	}
}

// New create a configurator
func New() *Configurator {
	return &Configurator{
		Bindings: struct {
			Transformers *Items
			Sorters      *Items
		}{
			&Items{},
			&Items{},
		},
	}
}

// findCategory retrieve a given category an return an error if category doesn't exist
func (c *Configurator) findCategory(category *string) (*Items, error) {
	switch *category {
	default:
		return nil, fmt.Errorf("category %s doesn't exist", *category)
	case "bindings.transformers":
		return c.Bindings.Transformers, nil
	case "bindings.sorters":
		return c.Bindings.Sorters, nil
	}
}

// Get from a category a value associated with a key
func (c *Configurator) Get(category *string, key *string) (string, error) {
	catConfig, err := c.findCategory(category)

	if err != nil {
		return "", err
	}

	if value, ok := (*catConfig)[*key]; ok == true {
		return value, nil
	}

	return "", fmt.Errorf("%s key doesn't exist", *key)
}

// Set associate a value to a key in a defined category
func (c *Configurator) Set(category *string, key *string, value *string) error {
	catConfig, err := c.findCategory(category)

	if err != nil {
		return err
	}

	(*catConfig)[*key] = *value

	return nil
}

// Delete from a category a given key
func (c *Configurator) Delete(category *string, key *string) error {
	catConfig, err := c.findCategory(category)

	if err != nil {
		return err
	}

	if _, ok := (*catConfig)[*key]; ok == false {
		return nil
	}

	delete((*catConfig), *key)

	return nil
}
