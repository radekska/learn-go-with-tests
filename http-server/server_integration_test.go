package main

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

func TestRecordWinsAndRetrievesThem(t *testing.T) {
	player := "Pepper"
	db, _ := GetDatabase()

	t.Run("test handles requests one by one", func(t *testing.T) {
		clearTable(t, db)

		store := NewPostgresPlayerStore(db)
		server := PlayerServer{store: store}

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)

		assert.Equal(t, "3", response.Body.String())

	})

	t.Run("test handles multiple score reads & writes at once", func(t *testing.T) {
		clearTable(t, db)

		store := NewPostgresPlayerStore(db)
		server := PlayerServer{store: store}
		readsAndWrites := 100

		var wg sync.WaitGroup
		wg.Add(readsAndWrites)

		for i := 0; i < readsAndWrites; i++ {
			go func() {
				server.ServeHTTP(httptest.NewRecorder(), newGetScoreRequest(player))
				server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
				wg.Done()
			}()
		}

		wg.Wait()

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assert.Equal(t, fmt.Sprintf("%d", readsAndWrites), response.Body.String())
	})

}
