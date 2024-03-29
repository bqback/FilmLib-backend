package middleware

import (
	"context"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/utils"
	"net/http"
	"strconv"
)

func ExtractID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		funcName := "ExtractID"

		logger, requestID, err := utils.GetLoggerAndID(r.Context())
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		idString := r.PathValue("id")

		if idString == "" {
			logger.DebugFmt("couldn't parse ID from URL", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}
		rCtx := context.WithValue(r.Context(), dto.IDKey, id)
		logger.DebugFmt("Stored in context", requestID, funcName, nodeName)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
