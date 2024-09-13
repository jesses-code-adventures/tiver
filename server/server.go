package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jesses-code-adventures/tiver/model"
)

type Server struct {
	Context *context.Context
	Conn    *pgx.Conn
	Logger  *slog.Logger
	Mux     *http.ServeMux
	Queries *model.Queries
}

func NewServer() (s Server) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	server := Server{
		Context: nil,
		Conn:    nil,
		Logger:  logger,
		Mux:     http.NewServeMux(),
		Queries: nil,
	}
	return server
}

func (s Server) ListenAndServe() (err error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, fmt.Sprintf("user=%s dbname=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE")))
	if err != nil {
		return
	}
	defer conn.Close(ctx)
	queries := model.New(conn)
	s.Queries = queries
	s.Context = &ctx
	s.Conn = conn
	s.registerHandlers()
	log.Print("successfully connected to db")
	err = http.ListenAndServe(os.Getenv("PORT"), s.Mux)
	if err != nil {
		return
	}
	return
}
