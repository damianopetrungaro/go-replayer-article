package user

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Storage map[string]User
	mu      sync.Mutex
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	id := chi.URLParam(r, "id")

	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.Storage[id]
	if !ok {
		http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, ErrGet.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) Save(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, ErrSave.Error(), http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Storage[u.ID] = u
}
