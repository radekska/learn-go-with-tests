package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary(t *testing.T) {
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
