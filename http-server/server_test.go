package poker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	assert.Equal(t, want, got, fmt.Sprintf("did not get correct status code, got %d, want %d", got, want))
}

func TestGetsScore(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{"Pepper": 20, "Floyd": 10}}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assert.Equal(t, "20", response.Body.String())
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assert.Equal(t, "10", response.Body.String())
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Missing")

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{}}
	server := NewPlayerServer(&store)

	t.Run("it returns created on POST", func(t *testing.T) {
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusCreated)
		AssertPlayerWin(t, &store, "Pepper")
	})
}

const jsonContentType = "application/json"

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"John", 14},
		}
		store := StubPlayerStore{scores: nil, winCalls: nil, league: wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assert.Equal(t, wantedLeague, got)
	})
}
func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	assert.Equalf(t, want, response.Result().Header.Get("content-type"), "response did not have content-type of %s, got %v", want, response.Result().Header)
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	league, _ = NewLeague(body)
	return league
}
