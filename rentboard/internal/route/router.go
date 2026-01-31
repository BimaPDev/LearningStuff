package route

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

type HealthResponse struct {
	OK bool `json:"ok"`
}
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/health", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(HealthResponse{OK: true})
	})

	r.Post("/auth/register", func(w http.ResponseWriter, req *http.Request) {
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if user.Name == "" || user.Email == "" || user.Password == "" {
			http.Error(w, "missing required fields", http.StatusBadRequest)
			return
		}

		if !strings.Contains(user.Email, "@") {
			http.Error(w, "invalid email format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "user registered successfully",
			"user": map[string]string{
				"id":    "12345",
				"name":  user.Name,
				"email": user.Email,
			},
		})
	})

	return r
}
