package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRecordWinsAndRetrievesThem(t *testing.T) {

	player := "Pepper"

	t.Run("test handles requests one by one", func(t *testing.T) {
		store := NewInMemoryPlayerStore()
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
		store := NewInMemoryPlayerStore()
		server := PlayerServer{store: store}
		readsAndWrites := 10000

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
