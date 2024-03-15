package middleware

import (
	"filmlib/internal/utils"
	"log"
	"net/http"
)

func JsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, err := utils.GetReqLogger(r.Context())
		if err != nil {
			log.Panic(err.Error())
		}
		requestID, err := utils.GetReqID(r.Context())
		if err != nil {
			log.Panic(err.Error())
		}
		funcName := "JsonHeader"

		w.Header().Set("Content-Type", "application/json")
		logger.DebugFmt("Content type header set", requestID, funcName, nodeName)

		next.ServeHTTP(w, r)
	})
}
