package v1

import (
	"net/http"
)

func (v *Router) Endpoints() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	return http.StripPrefix("/api", mux)
}
