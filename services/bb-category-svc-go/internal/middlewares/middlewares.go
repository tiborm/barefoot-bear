package middlewares

import "net/http"

func AddMiddlewares(next http.Handler) http.Handler {
	return LoggerMiddleware(next)
}