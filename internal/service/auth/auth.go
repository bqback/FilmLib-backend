package auth

import (
	"context"
	"filmlib/internal/auth"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/storage"
	"filmlib/internal/utils"
)

type AuthService struct {
	am *auth.AuthManager
	as storage.IAuthStorage
}

func NewAuthService(authStorage storage.IAuthStorage, manager *auth.AuthManager) *AuthService {
	return &AuthService{
		am: manager,
		as: authStorage,
	}
}

func (s *AuthService) Authenticate(ctx context.Context, info dto.LoginInfo) (*dto.JWT, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "Authenticate"

	ok, user, err := s.as.Auth(ctx, info)
	if !ok {
		return nil, err
	}
	logger.DebugFmt("Login info correct", requestID, funcName, "service")
	tokenString, err := s.am.GenerateToken(info, user)
	if err != nil {
		logger.DebugFmt("Token not generated "+err.Error(), requestID, funcName, "service")
		return nil, err
	}
	return &dto.JWT{Token: tokenString}, nil
}
