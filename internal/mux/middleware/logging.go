package middleware

import (
	"context"
	"filmlib/internal/apperrors"
	"filmlib/internal/logging"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/utils"
	"net/http"
)

func NewLogger(logger logging.ILogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			funcName := "SetLogger"
			requestLogger := logger
			reqID, err := utils.GetReqID(r.Context())
			if err != nil {
				requestLogger.Error(err.Error())
				apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
				return
			}

			rCtx := context.WithValue(r.Context(), dto.LoggerKey, requestLogger)
			requestLogger.DebugFmt("Added logger to context", reqID, funcName, nodeName)

			next.ServeHTTP(w, r.WithContext(rCtx))
		})
	}
}
