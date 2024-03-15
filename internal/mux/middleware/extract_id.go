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

		idString := r.PathValue("id")
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
		if idString == "" {
			logger.DebugFmt("couldn't parse ID from URL", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}
		id, err := strconv.Atoi(idString)
		if err != nil {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}
		rCtx := context.WithValue(r.Context(), dto.IDKey, id)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
