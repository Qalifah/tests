package player

import (
	"fmt"
	"net/http"
	"strings"
)

// Server controls data sent to the user
type Server struct {
	Store Collection
}


// ServeHTTP handles http request and gives a response
func(s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		s.processWin(w, player)
	case http.MethodGet:
		s.showScore(w, player)
	}
}

func(s *Server) showScore(w http.ResponseWriter, player string) {
	score := s.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func(s *Server) processWin(w http.ResponseWriter, player string) {
	s.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// Collection provides access to a collection of player's scores
type Collection interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// StubCollection stores players scores
type StubCollection struct {
	Scores map[string]int
	WinCalls	[]string
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

// NewInMemoryStore creates new in-memory store
func NewInMemoryStore() *InMemoryStore {
    return &InMemoryStore{map[string]int{}}
}

// InMemoryStore simulates an in-memory storage
type InMemoryStore struct {
    store map[string]int
}

// RecordWin returns
func (i *InMemoryStore) RecordWin(name string) {
    i.store[name]++
}

// GetPlayerScore returns  a player score from our in-memory storage
func (i *InMemoryStore) GetPlayerScore(name string) int {
    return i.store[name]
}