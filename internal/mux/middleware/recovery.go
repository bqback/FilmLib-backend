package middleware

import (
	"filmlib/internal/apperrors"
	"filmlib/internal/utils"
	"fmt"
	"log"
	"net/http"
)

const nodeName = "middleware"

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcvr := recover(); rcvr != nil {
				logger := *utils.GetReqLogger(r.Context())
				if logger == nil {
					log.Fatal("Logger missing from context")
					apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
				}
				logger.Error("*************** PANIC ***************")
				logger.Error(fmt.Sprintf("Recovered from panic %v", rcvr))

				apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)

				logger.Error("*************** CONTINUING ***************")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
