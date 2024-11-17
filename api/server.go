package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewServer() *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/users", PostUsersHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
