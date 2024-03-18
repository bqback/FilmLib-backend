package middleware

import (
	"filmlib/internal/apperrors"
	"filmlib/internal/auth"
	"filmlib/internal/utils"
	"net/http"
)

type AuthMiddleware struct {
	manager *auth.AuthManager
}

func NewAuthMiddleware(manager *auth.AuthManager) AuthMiddleware {
	return AuthMiddleware{manager: manager}
}

func (am AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, requestID, err := utils.GetLoggerAndID(r.Context())
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		funcName := "Auth"

		token := r.Header.Get("Authorization")
		if token == "" {
			logger.Error("Unauthorized access")
			apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
			return
		}
		logger.DebugFmt("Token found", requestID, funcName, nodeName)
		err = am.manager.ValidateToken(token)
		if err != nil {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
			logger.Error("Invalid authorization")
			apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
			return
		}
		logger.DebugFmt("Token validated", requestID, funcName, nodeName)

		next.ServeHTTP(w, r)
	})
}

func (am AuthMiddleware) CheckPerms(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, _, err := utils.GetLoggerAndID(r.Context())
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		token := r.Header.Get("Authorization")
		// if token == "" {
		// 	logger.Error("unauthorized access")
		// 	apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
		// 	return
		// }
		// err = auth.ValidateToken(token)
		// if err != nil {
		// 	logger.Error("invalid authorization")
		// 	apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
		// 	return
		// }
		hasRights, err := am.manager.CheckPerms(token)
		if err != nil {
			logger.Error("failed to check user permissions")
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}
		if !hasRights {
			logger.Error("insufficient rights")
			apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
