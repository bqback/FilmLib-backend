package middleware

import "net/http"

const nodeName = "middleware"

type MiddlewareFunc func(http.Handler) http.Handler

func Stack(handler http.Handler, middlewares ...MiddlewareFunc) http.Handler {
	for _, mw := range middlewares {
		handler = mw(handler)
	}
	return handler
}
