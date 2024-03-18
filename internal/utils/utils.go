package utils

import (
	"context"
	"filmlib/internal/apperrors"
	"filmlib/internal/logging"
	"filmlib/internal/pkg/dto"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetReqLogger(ctx context.Context) (logging.ILogger, error) {
	if ctx == nil {
		return nil, apperrors.ErrNilContext
	}
	if logger, ok := ctx.Value(dto.LoggerKey).(logging.ILogger); ok {
		return logger, nil
	}
	return nil, apperrors.ErrLoggerMissing
}

func GetReqID(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", apperrors.ErrNilContext
	}
	if reqID, ok := ctx.Value(dto.RequestIDKey).(string); ok {
		return reqID, nil
	}
	return "", apperrors.ErrRequestIDMissing
}

func GetIDParam(ctx context.Context) (uint64, error) {
	if ctx == nil {
		return 0, apperrors.ErrNilContext
	}
	if id, ok := ctx.Value(dto.IDKey).(uint64); ok {
		return id, nil
	}
	return 0, apperrors.ErrURLParamMissing
}

func GetSortOpts(ctx context.Context) (dto.SortOptions, error) {
	if ctx == nil {
		return dto.SortOptions{}, apperrors.ErrNilContext
	}
	if opts, ok := ctx.Value(dto.SortOptionsKey).(dto.SortOptions); ok {
		return opts, nil
	}
	return dto.SortOptions{}, apperrors.ErrSortOptionsMissing
}

func GetSearchQuery(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", apperrors.ErrNilContext
	}
	if query, ok := ctx.Value(dto.SearchTermKey).(string); ok {
		return query, nil
	}
	return "", apperrors.ErrSortOptionsMissing
}

func GetLoggerAndID(ctx context.Context) (logging.ILogger, string, error) {
	logger, err := GetReqLogger(ctx)
	if err != nil {
		log.Println(apperrors.ErrLoggerMissing)
		return nil, "", err
	}
	requestID, err := GetReqID(ctx)
	if err != nil {
		logger.Error(err.Error())
		return nil, "", err
	}
	return logger, requestID, nil
}

func HashFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePasswords(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
