package http_server

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func addScoreRecord(t *testing.T, db *sql.DB, name string, value int) {
	t.Helper()
	_, err := db.Exec("INSERT INTO players VALUES($1, $2)", name, value)
	if err != nil {
		t.Fatalf("could not store test record for player %s with %d - %v", name, value, err)
	}
}

func TestDatabaseStore(t *testing.T) {
	db, err := GetDatabase()
	if err != nil {
		t.Fatalf("failed to obtain DB connection: %v", err)
	}
	clearTable(t, db)
	store := NewPostgresPlayerStore(db)

	t.Run("get sorted league", func(t *testing.T) {
		defer clearTable(t, db)

		addScoreRecord(t, db, "Chris", 20)
		addScoreRecord(t, db, "Mike", 76)
		addScoreRecord(t, db, "Jake", 5)

		got := store.GetLeague()

		want := League{
			{Name: "Mike", Wins: 76},
			{Name: "Chris", Wins: 20},
			{Name: "Jake", Wins: 5},
		}

		assert.Equal(t, want, got)

	})

	t.Run("get player score", func(t *testing.T) {
		defer clearTable(t, db)

		addScoreRecord(t, db, "Chris", 20)

		got := store.GetPlayerScore("Chris")
		want := 20

		assert.Equal(t, want, got)
	})

	t.Run("store win for existing player", func(t *testing.T) {
		defer clearTable(t, db)

		addScoreRecord(t, db, "Chris", 20)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 21

		assert.Equal(t, want, got)
	})

	t.Run("store win for non-existing player", func(t *testing.T) {
		defer clearTable(t, db)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 1

		assert.Equal(t, want, got)
	})

}
