package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionarySearch(t *testing.T) {
	dictionary := Dictionary{"test": "just some random string for testing"}

	t.Run("exists", func(t *testing.T) {
		got, err := dictionary.Search("test")

		assert.NoError(t, err)
		assert.Equal(t, "just some random string for testing", got)
	})

	t.Run("not exists", func(t *testing.T) {
		_, err := dictionary.Search("vanilla")

		assert.Error(t, err)
		assert.Equal(t, ErrNotFound, err)

	})

}

func TestDictionaryAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "random string for test"

		err := dictionary.Add(word, definition)
		assert.NoError(t, err)

		got, err := dictionary.Search("test")
		assert.NoError(t, err)
		assert.Equal(t, definition, got)
	})

	t.Run("word exists", func(t *testing.T) {
		definition := "i already exists"
		dictionary := Dictionary{"test": definition}
		err := dictionary.Add("test", "trying to overwrite")

		got, _ := dictionary.Search("test")

		assert.Error(t, err)
		assert.Equal(t, ErrAlreadyExists, err)
		assert.Equal(t, definition, got)
	})

}

func TestDictionaryUpdate(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		word := "test"
		definition := "old value"
		dictionary := Dictionary{word: definition}
		newDefinition := "new value"

		err := dictionary.Update(word, newDefinition)
		assert.NoError(t, err)

		got, err := dictionary.Search(word)
		assert.NoError(t, err)
		assert.Equal(t, newDefinition, got)
	})

	t.Run("not exists", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("test", "i do not exist")
		assert.Error(t, err)
		assert.Equal(t, ErrWordDoesNotExist, err)
	})

}
