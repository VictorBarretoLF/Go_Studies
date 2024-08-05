package main

import (
	"context"
	"errors"
	"fmt"
	"go-web-socket/internal/api"
	"go-web-socket/internal/store/pgstore"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("[ERROR]", "[WSRS MAIN LOADING ENV]")
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("WSRS_DATABASE_HOST"),
			os.Getenv("WSRS_DATABASE_PORT"),
			os.Getenv("WSRS_DATABASE_USER"),
			os.Getenv("WSRS_DATABASE_PASSWORD"),
			os.Getenv("WSRS_DATABASE_NAME"),
		))

	if err != nil {
		fmt.Println("[ERROR]", "[CONNECTION POOL]")
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pgstore.New(pool))

	go func() {
		err := http.ListenAndServe(":8080", handler)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
