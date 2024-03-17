package middleware

import (
	"context"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/utils"
	"net/http"
	"strings"
)

func ExtractQuery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		funcName := "ExtractQuery"

		logger, requestID, err := utils.GetLoggerAndID(r.Context())
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		searchString := strings.ToLower(r.URL.Query().Get("query"))

		if searchString == "" {
			logger.DebugFmt("empty search", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}
		rCtx := context.WithValue(r.Context(), dto.SearchTermKey, searchString)
		logger.DebugFmt("Stored in context", requestID, funcName, nodeName)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
