package auth

import (
	"filmlib/internal/apperrors"
	"filmlib/internal/config"
	"filmlib/internal/pkg/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthManager struct {
	secret   []byte
	lifetime time.Duration
}

type jwtClaim struct {
	ID      uint64
	Name    string
	IsAdmin bool
	jwt.RegisteredClaims
}

func NewManager(config *config.JWTConfig) *AuthManager {
	return &AuthManager{secret: []byte(config.Secret), lifetime: config.Lifetime}
}

func (am *AuthManager) GenerateToken(info dto.LoginInfo, user *dto.DBUser) (string, error) {
	expiresAt := time.Now().Add(am.lifetime)
	claims := &jwtClaim{
		ID:      user.ID,
		Name:    user.Login,
		IsAdmin: user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(am.secret)
}

func (am *AuthManager) ValidateToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(am.secret), nil
		},
	)
	if err != nil {
		return err
	}

	claims, ok := parsedToken.Claims.(*jwtClaim)
	if !ok {
		return apperrors.ErrCouldNotParseClaims
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		return apperrors.ErrTokenExpired
	}
	if claims.IssuedAt.After(time.Now().Local()) {
		return apperrors.ErrInvalidIssuedTime
	}

	return nil
}

func (am *AuthManager) CheckPerms(token string) (bool, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(am.secret), nil
		},
	)
	if err != nil {
		return false, err
	}

	claims, ok := parsedToken.Claims.(*jwtClaim)
	if !ok {
		return false, apperrors.ErrCouldNotParseClaims
	}

	return claims.IsAdmin, nil
}
