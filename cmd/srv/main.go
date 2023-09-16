package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/go-replayer-article/user"
)

func main() {
	srv := user.Server{Storage: map[string]user.User{}}

	r := chi.NewRouter()

	r.Get("/users/{id}", srv.Get)
	r.Post("/users", srv.Save)

	log.Println("http server started running on port 8080")
	if err := http.ListenAndServe(":8080", r); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("http server stopped: %s", err)
	}
}
