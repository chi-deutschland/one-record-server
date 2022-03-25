package http

import "net/http"

type ContextHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}
