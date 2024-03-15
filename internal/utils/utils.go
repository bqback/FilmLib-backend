package utils

import (
	"context"
	"filmlib/internal/apperrors"
	"filmlib/internal/logging"
	"filmlib/internal/pkg/dto"
	"log"
)

func GetReqLogger(ctx context.Context) (logging.ILogger, error) {
	if ctx == nil {
		return nil, apperrors.ErrNilContext
	}
	if logger, ok := ctx.Value(dto.LoggerKey).(logging.ILogger); ok {
		return logger, nil
	}
	return nil, apperrors.ErrLoggerMissingFromContext
}

func GetReqID(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", apperrors.ErrNilContext
	}
	if reqID, ok := ctx.Value(dto.RequestIDKey).(string); ok {
		return reqID, nil
	}
	return "", apperrors.ErrRequestIDMissingFromContext
}

func GetLoggerAndID(ctx context.Context) (logging.ILogger, string, error) {
	logger, err := GetReqLogger(ctx)
	if err != nil {
		log.Println(apperrors.ErrLoggerMissingFromContext)
		return nil, "", err
	}
	requestID, err := GetReqID(ctx)
	if err != nil {
		logger.Error(err.Error())
		return nil, "", err
	}
	return logger, requestID, nil
}
