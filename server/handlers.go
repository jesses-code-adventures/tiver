package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesses-code-adventures/tiver/request"
)

func (s Server) startGame(w http.ResponseWriter, r *http.Request) {
	log.Printf("startGame got request %v", r.Header)
	game, err := s.Queries.CreateGame(*s.Context)
	if err != nil {
		log.Print("error creating a game")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoded := request.GameFromDbModel(game)
	asJson, err := json.Marshal(encoded)
	if err != nil {
		return
	}
	log.Printf("returning encoded: %s", asJson)
	w.Write(asJson)
}

func (s Server) passRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("passRequest got request. headers = %v, body = %v", r.Header, r.Body)
	db, err := s.Conn.Begin(*s.Context)
	if err != nil {
		log.Printf("error beginning transaction: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tx := s.Queries.WithTx(db)
	createRequestParams, err := request.RequestDbParamsFromSenderBody(r)
	if err != nil {
		log.Fatal("error getting db params")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	modelRequest, err := tx.CreateRequest(*s.Context, createRequestParams)
	if err != nil {
		log.Print("error creating request in db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	req := request.RequestFromDbModel(modelRequest)
	asJson, err := json.Marshal(req)
	if err != nil {
		log.Fatal("error marshalling to json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("returning encoded: %s", asJson)
	w.Write(asJson)
}

func (s Server) registerHandlers() {
	s.Mux.HandleFunc("/", s.startGame)
	s.Mux.HandleFunc("/request", s.passRequest)
}
