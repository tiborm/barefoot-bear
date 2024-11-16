package middlewares

import (
	"context"
	"net/http"

	"github.com/tiborm/barefoot-bear/pkg/logger"
)

type key string

const loggerKey key = "logger"


// LoggerMiddleware injects the logger into the request context.
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), loggerKey, logger.GetLogger())
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetLoggerFromContext retrieves the logger from the context.
func GetLoggerFromContext(ctx context.Context) *logger.Logger {
    if logger, ok := ctx.Value(loggerKey).(*logger.Logger); ok {
        return logger
    }
    return &logger.Logger{} // Fallback to a new logger if not found
}