// A standalone package for sending requests to the application while testing.
package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"math/rand"

	"github.com/jesses-code-adventures/tiver/request"
)

type Sender struct {
	Logger         slog.Logger
	Scheme         string
	Host           string
	ReceiverScheme string
	ReceiverHost   string
	ReceiverPort   *string
}

func NewHttpSender() Sender {
	scheme := os.Getenv("SENDER_SCHEME")
	host := os.Getenv("SENDER_HOST")
	port := os.Getenv("SENDER_PORT")
	if port != "" {
		host += port
	}
	receiverScheme := os.Getenv("SCHEME")
	receiverHost := os.Getenv("HOST")
	receiverPort := os.Getenv("PORT")
	if receiverPort != "" {
		receiverHost += receiverPort
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return Sender{Scheme: scheme, Host: host, ReceiverScheme: receiverScheme, ReceiverHost: receiverHost, ReceiverPort: &receiverPort}
}

func (s *Sender) url(endpoint string) (u string) {
	joined, err := url.Parse(fmt.Sprintf("%s%s%s", s.Scheme, s.Host, endpoint))
	if err != nil {
		panic(fmt.Sprintf("Error joining url.\n\nScheme: %s\nHost%s \nEndpoint: %s\n\nError: %s", s.Scheme, s.Host, endpoint, err.Error()))
	}
	u = joined.String()
	return
}

func (s *Sender) urlForReceiver(endpoint string) (u string) {
	joined, err := url.Parse(fmt.Sprintf("%s%s%s", s.ReceiverScheme, s.ReceiverHost, endpoint))
	if err != nil {
		panic(fmt.Sprintf("Error joining url.\n\nScheme: %s\nHost%s \nEndpoint: %s\n\nError: %s", s.Scheme, s.Host, endpoint, err.Error()))
	}
	u = joined.String()
	return
}

func (s *Sender) startGame() (game request.Game, err error) {
	resp, err := http.Get(s.urlForReceiver("/"))
	if err != nil {
		log.Printf("error in startGame: %s", err.Error())
		return
	}
	return request.GameFromResponse(resp)
}

func (s *Sender) sendRequest(r request.IncomingRequest) (err error) {
	encoded, err := json.Marshal(r)
	if err != nil {
		return
	}
	reader := bytes.NewBuffer(encoded)
	log.Printf("reader bytes = %s", reader)
	resp, err := http.Post(s.urlForReceiver("/request"), "application/json", reader)
	if err != nil {
		return
	}
	log.Printf("status code: %s", resp.Status)
	return
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }
	// defer resp.Body.Close()
	// log.Printf("resp status code: %v")
}

func (s *Sender) SendRequests() (err error) {
	game, err := s.startGame()
	if err != nil {
		return
	}
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(3-1)+1) * time.Millisecond * 10)
			req := request.IncomingRequestDummy(game.Id)
			if err = s.sendRequest(req); err != nil {
				log.Fatal(err)
			}
		}
	}()
	select {}
}
