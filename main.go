package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgtype"
)

func run() (err error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=jesse dbname=tiver sslmode=disable")
	if err != nil {
		return
	}
	defer conn.Close(ctx)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	log.Print("successfully connected to db")
	for {

	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
