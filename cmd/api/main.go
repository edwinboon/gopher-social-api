package main

import (
	"log"

	"github.com/edwinboon/gopher-social-api/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetEnv("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.serve(mux))
}
