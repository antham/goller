package configurator

import (
	"testing"
)

func TestGetFromUnexistingCategory(t *testing.T) {
	c := New()
	category := "test"
	key := "whatever"
	_, err := c.Get(&category, &key)

	if err.Error() != "category test doesn't exist" {
		t.Error("Must throw an error")
	}
}

func TestGetUnexistingKey(t *testing.T) {
	c := New()
	category := "bindings.transformers"
	key := "whatever"
	_, err := c.Get(&category, &key)

	if err.Error() != "whatever key doesn't exist" {
		t.Error("Must throw an error")
	}
}

func TestGetExistingKey(t *testing.T) {
	c := New()
	c.Bindings.Sorters = &Items{
		"hello": "world",
	}
	category := "bindings.sorters"
	key := "hello"
	value, _ := c.Get(&category, &key)

	if value != "world" {
		t.Error("Must retrieve value world")
	}
}

func TestSet(t *testing.T) {
	c := New()

	category := "bindings.transformers"
	key := "hello"
	value := "world"

	c.Set(&category, &key, &value)

	if (*c.Bindings.Transformers)["hello"] != "world" {
		t.Error("Must retrieve value world")
	}
}

func TestSetFromUnexistingCategory(t *testing.T) {
	c := New()

	category := "test"
	key := "hello"
	value := "world"

	err := c.Set(&category, &key, &value)

	if err.Error() != "category test doesn't exist" {
		t.Error("Must throw an error")
	}
}

func TestDelete(t *testing.T) {
	c := New()
	c.Bindings.Transformers = &Items{
		"hello": "world",
	}

	category := "bindings.transformers"
	key := "hello"

	c.Delete(&category, &key)

	if _, ok := (*c.Bindings.Transformers)["hello"]; !ok {
		t.Error("Must delete key hello")
	}
}

func TestDeleteFromUnexistingCategory(t *testing.T) {
	c := New()

	category := "test"
	key := "hello"

	err := c.Delete(&category, &key)

	if err.Error() != "category test doesn't exist" {
		t.Error("Must throw an error")
	}
}

func TestDeleteUnexistingKey(t *testing.T) {
	c := New()

	category := "bindings.transformers"
	key := "hell"

	err := c.Delete(&category, &key)

	if err != nil {
		t.Error("Must not throw an error")
	}
}
