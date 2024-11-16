package routes

import (
	"net/http"

	"github.com/tiborm/barefoot-bear/bb-catalog-svc-go/internal/handlers"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HandleMain)
    mux.HandleFunc("/login", handlers.HandleGoogleLogin)
    mux.HandleFunc("/callback", handlers.HandleGoogleCallback)
}