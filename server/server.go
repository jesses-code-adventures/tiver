package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Server struct {
	Context *context.Context
	Logger  *slog.Logger
	Mux     *http.ServeMux
}

func NewServer() (s Server) {
	err := godotenv.Load(".env", ".env.secret")
	if err != nil {
		panic("no env file found")
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=jesse dbname=tiver sslmode=disable")
	if err != nil {
		return
	}
	defer conn.Close(ctx)
	log.Print("successfully connected to db")
	return Server{
		Context: &ctx,
		Logger:  logger,
		Mux:     http.NewServeMux(),
	}
}

func (s Server) ListenAndServe() (err error) {
	err = http.ListenAndServe(os.Getenv("PORT"), s.Mux)
	if err != nil {
		return
	}
	return
}
