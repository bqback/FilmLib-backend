package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrNilContext = errors.New("context is nil")
)

var (
	ErrInvalidLoggingLevel      = errors.New("invalid logging level")
	ErrLoggerMissingFromContext = errors.New("logger missing from context")
)

var (
	ErrRequestIDMissingFromContext = errors.New("request ID is missing from context")
)

var (
	ErrCouldNotParseURLParam = errors.New("failed to parse URL params")
)

var (
	ErrEnvNotFound       = errors.New("unable to load .env file")
	ErrDatabasePWMissing = errors.New("database password is missing from env")
)

var (
	ErrCouldNotBuildQuery       = errors.New("failed to build SQL query")
	ErrCouldNotBeginTransaction = errors.New("failed to start DB transaction")
	ErrCouldNotRollback         = errors.New("failed to roll back after a failed query")
	ErrCouldNotCommit           = errors.New("failed to commit DB transaction changes")
)

var (
	ErrActorNotCreated = errors.New("failed to insert actor into database")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var BadRequestResponse = ErrorResponse{
	Code:    http.StatusBadRequest,
	Message: "Bad request",
}

var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "Internal error",
}

func ReturnError(err ErrorResponse, w http.ResponseWriter, r *http.Request) {
	// rCtx := r.Context()
	// logger := *utils.GetReqLogger(rCtx)
	// requestID := chimw.GetReqID(rCtx)
	// funcName := "ReturnError"

	w.WriteHeader(err.Code)
	// logger.DebugFmt(fmt.Sprintf("Wrote error code %v", err.Code), requestID, funcName, nodeName)
	_, _ = w.Write([]byte(err.Message))
	// logger.DebugFmt(fmt.Sprintf("Wrote message %v", err.Code), requestID, funcName, nodeName)
	r.Body.Close()
	// logger.DebugFmt("Request body closed", requestID, funcName, nodeName)
}
