package middleware

import (
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// requestID := uuid.New().String()
		// rCtx := context.WithValue(r.Context(), dto.RequestIDKey, requestID)

		next.ServeHTTP(w, r)
	})
}
