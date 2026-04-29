package main

import (
	"log"
	"time"

	"github.com/edwinboon/gopher-social-api/internal/db"
	"github.com/edwinboon/gopher-social-api/internal/env"
	"github.com/edwinboon/gopher-social-api/internal/store"
)

func main() {
	addr := env.GetEnv("DB_DSN", "postgres://admin:admin@localhost/social?sslmode=disable")

	conn, err := db.New(addr, 3, 3, time.Minute*15)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStore(conn)
	db.Seed(store)
}

