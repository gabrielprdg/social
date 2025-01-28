package main

import (
	"github.com/gabrielprdg/social.git/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type dbConfig struct {
	address      string
	maxOpenConns string
	maxIdleConns string
	maxIdleTime  string
}

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr    string
	db      dbConfig
	env     string
	version string
}

func (app *application) mount() http.Handler {
	// create router
	r := chi.NewRouter()

	// middleware to show all the requests to our server
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/users", app.healthCheckHandler)
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
			r.Route("/{postID}", func(r chi.Router) {
				r.Get("/", app.getPostByIdHandler)
			})
		})
	})
	return r
}

func (app *application) run(mux http.Handler) error {
	// request router and dispatch

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("listening on %s", app.config.addr)
	return server.ListenAndServe()
}
