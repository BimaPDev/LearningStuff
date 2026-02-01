package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/bcrypt"
)

type HealthResponse struct {
	OK bool `json:"ok"`
}
type UserRegister struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type StoreUser struct {
	ID           string
	Name         string
	Email        string
	PasswordHash []byte
}

var usersByEmail = map[string]StoreUser{}

func hashPassworrd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(pass, hpass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hpass), []byte(pass))
	if err != nil {
		return err
	}
	return nil
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
		var user UserRegister
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, ok := usersByEmail[user.Email]
		if ok {
			w.WriteHeader(http.StatusConflict)
			return
		}

		hashed, err := hashPassworrd(user.Password)
		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			return
		}

		usersByEmail[user.Email] = StoreUser{
			ID:           "12345", // (still hardcoded for now)
			Name:         user.Name,
			Email:        user.Email,
			PasswordHash: []byte(hashed),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(UserRegisterResponse{
			ID:    "12345",
			Name:  user.Name,
			Email: user.Email,
		})

		for email, u := range usersByEmail {
			fmt.Printf("stored user: email=%s id=%s name=%s password_hash=%s\n", email, u.ID, u.Name, u.PasswordHash)
		}

	})

	r.Post("/auth/login", func(w http.ResponseWriter, req *http.Request) {
		var login UserLogin
		err := json.NewDecoder(req.Body).Decode(&login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stored, ok := usersByEmail[login.Email]
		if !ok {
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
			return
		}
		if err := CheckPassword(login.Password, string(stored.PasswordHash)); err != nil {
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(UserLoginResponse{
			Name:  login.Name,
			Email: login.Email,
		},
		)
	})

	return r
}
