package middleware

import (
	"context"
	"filmlib/internal/pkg/dto"
	"net/http"

	"github.com/google/uuid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		rCtx := context.WithValue(r.Context(), dto.RequestIDKey, requestID)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
