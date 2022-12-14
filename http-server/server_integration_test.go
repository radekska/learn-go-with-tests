package poker

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func clearTable(t *testing.T, db *sql.DB) {
	t.Helper()
	_, err := db.Exec("DELETE FROM players")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecordWinsAndRetrieveThem(t *testing.T) {
	playerName := "Pepper"
	db, _ := GetDatabase()

	t.Run("test handles requests one by one", func(t *testing.T) {
		clearTable(t, db)

		store := NewPostgresPlayerStore(db)
		server := NewPlayerServer(store)

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(playerName))

		assertStatus(t, response.Code, http.StatusOK)

		assert.Equal(t, "3", response.Body.String())

	})

	t.Run("test handles multiple score reads & writes at once", func(t *testing.T) {
		clearTable(t, db)

		store := NewPostgresPlayerStore(db)
		server := NewPlayerServer(store)
		readsAndWrites := 100

		var wg sync.WaitGroup
		wg.Add(readsAndWrites)

		for i := 0; i < readsAndWrites; i++ {
			go func() {
				server.ServeHTTP(httptest.NewRecorder(), newGetScoreRequest(playerName))
				server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
				wg.Done()
			}()
		}

		wg.Wait()

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(playerName))

		assertStatus(t, response.Code, http.StatusOK)
		assert.Equal(t, fmt.Sprintf("%d", readsAndWrites), response.Body.String())
	})

}

func TestGetLeague(t *testing.T) {
	db, _ := GetDatabase()
	clearTable(t, db)

	playerName := "John"

	store := NewPostgresPlayerStore(db)
	server := NewPlayerServer(store)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newLeagueRequest())

	got := getLeagueFromResponse(t, response.Body)
	want := []Player{{playerName, 2}}

	assert.Equal(t, want, got)
	assertStatus(t, response.Code, http.StatusOK)
	assertContentType(t, response, jsonContentType)
}
