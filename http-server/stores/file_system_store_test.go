package stores

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"learn-go-with-tests/http-server/player"
)

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
	return tmpFile, removeFile
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err, "could not initialize file system store")

		got := store.GetLeague()

		want := player.League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assert.Equal(t, want, got)

		// read again
		got = store.GetLeague()
		assert.Equal(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err, "could not initialize file system store")

		got := store.GetPlayerScore("Chris")
		want := 33
		assert.Equal(t, want, got)
	})

	t.Run("store win for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err, "could not initialize file system store")

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assert.Equal(t, want, got)
	})

	t.Run("store win for non-existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err, "could not initialize file system store")

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 1
		assert.Equal(t, want, got)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assert.NoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assert.NoError(t, err)

		got := store.GetLeague()

		want := player.League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assert.Equal(t, want, got)

		got = store.GetLeague()
		assert.Equal(t, want, got)
	})
}
