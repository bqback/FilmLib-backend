package middleware

import (
	"filmlib/internal/apperrors"
	"filmlib/internal/utils"
	"fmt"
	"log"
	"net/http"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcvr := recover(); rcvr != nil {
				logger, err := utils.GetReqLogger(r.Context())
				if logger != nil {
					log.Fatal(err.Error())
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
