package handlers

import (
	"net/http"

	"github.com/tiborm/barefoot-bear/internal/middlewares"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLoggerFromContext(r.Context())
    logger.Request(r, "Handling main request")

	htmlIndex := `<html><body><a href="/login">Google Log In</a></body></html>`
	w.Write([]byte(htmlIndex))
}
