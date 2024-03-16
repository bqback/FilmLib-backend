package middleware

import (
	"context"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/utils"
	"net/http"
	"strconv"

	"log"
)

func ExtractID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		funcName := "ExtractID"

		logger, err := utils.GetReqLogger(r.Context())
		if err != nil {
			log.Println(err.Error())
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}
		requestID, err := utils.GetReqID(r.Context())
		if err != nil {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		idString := r.PathValue("id")

		if idString == "" {
			logger.DebugFmt("couldn't parse ID from URL", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}
		logger.DebugFmt("Extracted id as string: "+idString, requestID, funcName, nodeName)
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}
		logger.DebugFmt("Converted to uint64", requestID, funcName, nodeName)
		rCtx := context.WithValue(r.Context(), dto.IDKey, id)
		logger.DebugFmt("Stored in context", requestID, funcName, nodeName)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
