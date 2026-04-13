package main

import (
	"log"

	"github.com/edwinboon/gopher-social-api/internal/env"
	"github.com/edwinboon/gopher-social-api/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetEnv("ADDR", ":8080"),
	}

	store := store.NewStore(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.serve(mux))
}
