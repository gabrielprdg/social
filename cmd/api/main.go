package main

import (
	"github.com/gabrielprdg/social.git/internal/db"
	"github.com/gabrielprdg/social.git/internal/store"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const version = "0.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: ":" + os.Getenv("ADDR"),
		db: dbConfig{
			address:      os.Getenv("DB_ADDRESS"),
			maxOpenConns: os.Getenv("DB_MAX_OPEN_CONNS"),
			maxIdleConns: os.Getenv("DB_MAX_IDLE_CONNS"),
			maxIdleTime:  os.Getenv("DB_MAX_IDLE_TIME"),
		},
		env: os.Getenv("ENVIRONMENT"),
	}

	maxOpConns, err := strconv.Atoi(cfg.db.maxOpenConns)
	if err != nil {
		log.Fatal("Error converting max open connections to int")
	}

	maxIdConns, err := strconv.Atoi(cfg.db.maxIdleConns)
	if err != nil {
		log.Fatal("Error converting max open connections to int")
	}

	db, err := db.New(cfg.db.address, cfg.db.maxIdleTime, maxOpConns, maxIdConns)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
