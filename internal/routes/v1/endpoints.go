package v1

import (
	"net/http"
)

func (v *Router) Endpoints() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	mux.Handle("GET /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.GetUserByName(w, r)
	}))

	mux.Handle("POST /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.InsertUser(w, r)
	}))

	return http.StripPrefix("/api", mux)
}
