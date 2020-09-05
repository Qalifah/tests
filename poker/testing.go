package poker

import (
	"testing"
	"net/http/httptest"
	"reflect"
)

// StubCollection stores players scores
type StubCollection struct {
	Scores map[string]int
	WinCalls	[]string
	league 		League
}

// GetPlayerScore returns a player's score
func(s *StubCollection) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin stores a name in the collection
func(s *StubCollection) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

// GetLeague returns the league 
func(s *StubCollection) GetLeague() League {
	return s.league
}

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("didn't expect an error but got one, %v", err)
    }
}

func AssertScoreEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func AssertPlayerWin(t *testing.T, store *StubCollection, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("didn't store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}

func AssertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}