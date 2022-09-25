package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("say hello to people", func(t *testing.T) {
		assert.Equal(t, "Hello, Chris!", Hello("Chris", ""))
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		assert.Equal(t, "Hello, World!", Hello("", ""))
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		assert.Equal(t, "Hola, Elodie!", Hello("Elodie", "Spanish"))
	})

	t.Run("say hello in French", func(t *testing.T) {
		assert.Equal(t, "Bonjour, Juliet!", Hello("Juliet", "French"))
	})
}
