package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/go-chi/chi/v5"
)

func ExampleClient() {
	ctx := context.Background()

	srv := Server{Storage: map[string]User{}}

	r := chi.NewRouter()
	r.Get("/users/{id}", srv.Get)
	r.Post("/users", srv.Save)

	ts := httptest.NewServer(r)
	defer ts.Close()

	cl := Client{BaseClient: http.DefaultClient, BaseURL: ts.URL}

	found, err := cl.Get(ctx, "an-id")
	if !errors.Is(err, ErrNotFound) || found != nil {
		log.Fatalf("could not match the error of a non-existing user. Error: %v - User: %v", err, found)
	}

	u := &User{
		ID:        "an-id",
		Name:      "Luke",
		CreatedAt: time.Now(),
	}

	if err := cl.Save(ctx, u); err != nil {
		log.Fatalf("could not save the user. Error: %s", err)
	}

	found, err = cl.Get(ctx, "an-id")
	if err != nil {
		log.Fatalf("could not get the user. Error: %s", err)
	}

	fmt.Println(found.ID)
	fmt.Println(found.Name)

	// Output:
	// an-id
	// Luke
}
