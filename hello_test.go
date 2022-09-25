package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		assert.Equal(t, "Hello, Chris!", Hello("Chris"))
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		assert.Equal(t, "Hello, World!", Hello(""))
	})
}
