package utils

import (
	"context"
	"filmlib/internal/logging"
	"filmlib/internal/pkg/dto"
)

func GetReqLogger(ctx context.Context) *logging.ILogger {
	if ctx == nil {
		return nil
	}
	if logger, ok := ctx.Value(dto.LoggerKey).(logging.ILogger); ok {
		return &logger
	}
	return nil
}
