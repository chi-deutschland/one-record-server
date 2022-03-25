package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AuthHeaderMiddleware(k, v string) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			value := r.Header.Get(k)
			if v == value {
				h.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		})
	}
}
