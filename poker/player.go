package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Server controls data sent to the user
type Server struct {
	Store Collection
	http.Handler
}

//Player represents a user
type Player struct {
	Name string
	Wins	int
}

// NewServer initialises a new server
func NewServer(store Collection) *Server {
	s := new(Server)
	s.Store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playersHandler))

	s.Handler = router
	return s
}

func(s *Server) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(s.Store.GetLeague())
	w.WriteHeader(http.StatusOK)
}


func(s *Server) playersHandler(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/players/"):]

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
	GetLeague() League
}

