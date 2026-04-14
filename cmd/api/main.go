package main

import (
	"log"
	"time"

	"github.com/edwinboon/gopher-social-api/internal/db"
	"github.com/edwinboon/gopher-social-api/internal/env"
	"github.com/edwinboon/gopher-social-api/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetEnv("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetEnv("DB_ADDR", "postgres://admin:admin@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetEnvAsInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetEnvAsDuration("DB_MAX_IDLE_TIME", time.Minute*15),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("db connection pool established")

	store := store.NewStore(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.serve(mux))
}
