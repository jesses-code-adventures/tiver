package server

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/jesses-code-adventures/tiver/request"
)

func (s Server) startGame(w http.ResponseWriter, r *http.Request) {
	log.Printf("got request %v", r.Header)
	gameId, err := s.Queries.CreateGame(*s.Context)
	if err != nil {
		s.Logger.Error("error creating a game")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(gameId)
}

func (s Server) passRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("passRequest got request. headers = %v, body = %v", r.Header, r.Body)
}

func (s Server) registerHandlers() {
	s.Mux.HandleFunc("/", s.startGame)
}
