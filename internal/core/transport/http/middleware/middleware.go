package core_http_middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func ChainMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	if len(middlewares) == 0 {
		return handler
	}
	for idx := len(middlewares) - 1; idx >= 0; idx-- {
		handler = middlewares[idx](handler)
	}
	return handler
}
